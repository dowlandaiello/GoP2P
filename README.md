# GoP2P

A streamlined shard-based P2P networking stack built in Go.

[![Author](https://img.shields.io/badge/made%20by-Mitsuko%20Megumi-purple.svg?style=flat-round)](https://github.com/mitsukomegumi)
[![Project](https://img.shields.io/badge/project-GoP2P-blue.svg?style=flat-round)](https://github.com/mitsukomegumi/gop2p)
[![GoDoc](https://godoc.org/github.com/mitsukomegumi/gop2p?status.svg)](https://godoc.org/github.com/mitsukomegumi/gop2p)
[![Author](https://godoc.org/github.com/mitsukomegumi/gop2p?status.svg)](https://godoc.org/github.com/mitsukomegumi/gop2p)
[![Build Status](https://travis-ci.com/mitsukomegumi/GoP2P.svg?branch=master)](https://travis-ci.com/mitsukomegumi/GoP2P)
[![CircleCI](https://circleci.com/gh/mitsukomegumi/GoP2P.svg?style=svg)](https://circleci.com/gh/mitsukomegumi/GoP2P)
[![Go Report Card](https://goreportcard.com/badge/github.com/mitsukomegumi/gop2p)](https://goreportcard.com/report/github.com/mitsukomegumi/gop2p)
[![codecov](https://codecov.io/gh/mitsukomegumi/GoP2P/branch/master/graph/badge.svg)](https://codecov.io/gh/mitsukomegumi/GoP2P)

## Dependencies

All critical dependencies have been included in .vendor/. As of now, protoc (compile .proto files) is the only GoP2P dependency not shipped in .vendor or .vendor-new. Fortunately, protoc is only necessary for the proper use of GoP2P in the case that .pb.go files need to be compiled on runtime. In production environments where users cannot be prompted to install protoc, it is recommended that the necessary .pb.go files are compiled and provided by the developers of software using GoP2P before runtime.