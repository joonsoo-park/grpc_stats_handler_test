package main

import (
	"context"
	"flag"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/encoding/gzip"
	"google.golang.org/protobuf/proto"

	pb "dice/proto"
)

const (
	defaultName = `ashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jfashdjgk;hskldhjviop ewhfio ;hasiopfh s;aiodhfio ajwFIOP EHWFP AHWPRT89 H3RT89HY23PR9 8HSAPFUH IOPdfh aopihOPI FHAIPWFH IOPASEH FIOPAHSD FKL;Hiof haopwfhy	askdfhjakl;sdfhjl;kasdfjj akls;dasdgasdjkflxz;cklvjxcklz;jvkl;xzcjvkl;xcjvl;kxzcjklv;jzxcl;vkjzxcvkl;jzxklcvjl;ckxzjvkl;xzcjvl;kzjxcvl;kzxcjvklzcjxvz
	zkxcjv;lkjxzcklv;jxzclv;jf`
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// handler := MyHandler{}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		// grpc.WithStatsHandler(&handler),
	}

	conn, err := grpc.Dial(*addr, opts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	data := &pb.GreetRequest{Name: *name, Msg: "Hello", Age: 10}

	log.Println(proto.Size(data))

	r, err := c.Greet(ctx, data, grpc.UseCompressor(gzip.Name))
	// r, err := c.Greet(ctx, data)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetGreeting())
}
