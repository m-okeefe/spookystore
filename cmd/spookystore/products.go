// Copyright 2018 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"strconv"
	"time"

	"cloud.google.com/go/datastore"
	"cloud.google.com/go/trace"
	"golang.org/x/net/context"

	pb "github.com/m-okeefe/spookystore/internal/proto"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type productsDirectory struct {
	ds         *datastore.Client
	ctx        context.Context
	productIds []int64
}

type Product struct {
	K                    *datastore.Key `datastore:"__key__"`
	ID                   string         `datastore:"ID"`
	DisplayName          string         `datastore:"DisplayName"`
	PictureURL           string         `datastore:"PictureURL"`
	Cost                 float32        `datastore:"Cost"`
	Description          string         `datastore:"Description"`
	XXX_NoUnkeyedLiteral struct{}       `datastore:"XXX_NoUnkeyedLiteral"`
	XXX_unrecognized     []byte         `datastore:"XXX_unrecognized"`
	XXX_sizecache        int32          `datastore:"XXX_sizecache"`
}

type TransactionCounter struct {
	NumTransactions int32 `datastore:"NumTransactions"`
}

func (s *Server) GetNumTransactions(ctx context.Context, req *pb.GetNumTransactionsRequest) (*pb.NumTransactionsResponse, error) {
	var t TransactionCounter
	k := datastore.NameKey("TransactionCounter", "AllPurchases", nil)
	err := s.ds.Get(ctx, k, &t)
	if err != nil {
		log.WithField("error", err).Error("failed to get num transactions")
		return nil, errors.Wrap(err, "failed to query")
	}
	return &pb.NumTransactionsResponse{
		NumTransactions: t.NumTransactions,
	}, nil
}

func (s *Server) GetAllProducts(ctx context.Context, req *pb.GetAllProductsRequest) (*pb.GetAllProductsResponse, error) {
	span := trace.FromContext(ctx).NewChild("spookystoresvc/GetAllProducts")
	defer span.Finish()

	log := log.WithFields(logrus.Fields{
		"op": "GetAllProducts"})
	start := time.Now()
	defer func() {
		log.WithField("elapsed", time.Since(start).String()).Debug("completed request")
	}()
	log.Debug("received request")

	cs := span.NewChild("datastore/query/products")
	defer cs.Finish()

	var result []*pb.Product
	_, err := s.ds.GetAll(ctx, datastore.NewQuery("Product"), &result)
	if err != nil {
		log.WithField("error", err).Error("failed to query the datastore")
		return nil, errors.Wrap(err, "failed to getAll")
	}
	return &pb.GetAllProductsResponse{ProductList: result}, nil
}

func (s *Server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.Product, error) {
	var v Product
	parsed, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		return &pb.Product{}, nil
	}

	err = s.ds.Get(ctx, datastore.IDKey("Product", parsed, nil), &v)
	if err == datastore.ErrNoSuchEntity {
		log.Debug("product not found")
		return &pb.Product{}, nil
	} else if err != nil {
		log.WithField("error", err).Error("failed to query the datastore")
		return nil, errors.Wrap(err, "failed to query")
	}
	log.Debug("found product")
	return &pb.Product{
		ID:          fmt.Sprintf("%d", v.K.ID),
		DisplayName: v.DisplayName,
		PictureURL:  v.PictureURL,
		Cost:        v.Cost,
		Description: v.Description,
	}, nil
}

func (s *Server) AddProductToCart(ctx context.Context, req *pb.AddProductRequest) (*pb.AddProductResponse, error) {
	// get user
	userResp, err := s.GetUser(ctx, &pb.UserRequest{ID: req.UserID})
	if err != nil {
		return &pb.AddProductResponse{Success: false}, err
	}
	user := userResp.GetUser()

	// update cart
	items := user.Cart.GetItems()
	if items == nil {
		items = map[string]int32{}
	}
	if _, ok := items[req.ProductID]; ok {
		items[req.ProductID] = items[req.ProductID] + 1
	} else {
		items[req.ProductID] = 1
	}

	user.Cart.Items = items

	// put user
	parsed, err := strconv.ParseInt(user.ID, 10, 64)
	u := datastore.IDKey("User", parsed, nil)
	if _, err := s.ds.Put(ctx, u, user); err != nil {
		return &pb.AddProductResponse{Success: false}, err
	}
	return &pb.AddProductResponse{Success: true}, nil
}
