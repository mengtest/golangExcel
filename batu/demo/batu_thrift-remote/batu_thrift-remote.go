// Autogenerated by Thrift Compiler (0.9.3)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"batu/demo"
	"flag"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"math"
	"net"
	"net/url"
	"os"
	"strconv"
	"strings"
)

func Usage() {
	fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
	flag.PrintDefaults()
	fmt.Fprintln(os.Stderr, "\nFunctions:")
	fmt.Fprintln(os.Stderr, "  string CallBack(i64 callTime, i32 types,  paramMap)")
	fmt.Fprintln(os.Stderr, "  void put(Article newArticle)")
	fmt.Fprintln(os.Stderr)
	os.Exit(0)
}

func main() {
	flag.Usage = Usage
	var host string
	var port int
	var protocol string
	var urlString string
	var framed bool
	var useHttp bool
	var parsedUrl url.URL
	var trans thrift.TTransport
	_ = strconv.Atoi
	_ = math.Abs
	flag.Usage = Usage
	flag.StringVar(&host, "h", "localhost", "Specify host and port")
	flag.IntVar(&port, "p", 9090, "Specify port")
	flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
	flag.StringVar(&urlString, "u", "", "Specify the url")
	flag.BoolVar(&framed, "framed", false, "Use framed transport")
	flag.BoolVar(&useHttp, "http", false, "Use http")
	flag.Parse()

	if len(urlString) > 0 {
		parsedUrl, err := url.Parse(urlString)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
		host = parsedUrl.Host
		useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
	} else if useHttp {
		_, err := url.Parse(fmt.Sprint("http://", host, ":", port))
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
			flag.Usage()
		}
	}

	cmd := flag.Arg(0)
	var err error
	if useHttp {
		trans, err = thrift.NewTHttpClient(parsedUrl.String())
	} else {
		portStr := fmt.Sprint(port)
		if strings.Contains(host, ":") {
			host, portStr, err = net.SplitHostPort(host)
			if err != nil {
				fmt.Fprintln(os.Stderr, "error with host:", err)
				os.Exit(1)
			}
		}
		trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		if framed {
			trans = thrift.NewTFramedTransport(trans)
		}
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error creating transport", err)
		os.Exit(1)
	}
	defer trans.Close()
	var protocolFactory thrift.TProtocolFactory
	switch protocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
		break
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		break
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
		break
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
		Usage()
		os.Exit(1)
	}
	client := demo.NewBatuThriftClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "CallBack":
		if flag.NArg()-1 != 3 {
			fmt.Fprintln(os.Stderr, "CallBack requires 3 args")
			flag.Usage()
		}
		argvalue0, err8 := (strconv.ParseInt(flag.Arg(1), 10, 64))
		if err8 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		tmp1, err9 := (strconv.Atoi(flag.Arg(2)))
		if err9 != nil {
			Usage()
			return
		}
		argvalue1 := int32(tmp1)
		value1 := argvalue1
		arg10 := flag.Arg(3)
		mbTrans11 := thrift.NewTMemoryBufferLen(len(arg10))
		defer mbTrans11.Close()
		_, err12 := mbTrans11.WriteString(arg10)
		if err12 != nil {
			Usage()
			return
		}
		factory13 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt14 := factory13.GetProtocol(mbTrans11)
		containerStruct2 := demo.NewBatuThriftCallBackArgs()
		err15 := containerStruct2.ReadField3(jsProt14)
		if err15 != nil {
			Usage()
			return
		}
		argvalue2 := containerStruct2.ParamMap
		value2 := argvalue2
		fmt.Print(client.CallBack(value0, value1, value2))
		fmt.Print("\n")
		break
	case "put":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Put requires 1 args")
			flag.Usage()
		}
		arg16 := flag.Arg(1)
		mbTrans17 := thrift.NewTMemoryBufferLen(len(arg16))
		defer mbTrans17.Close()
		_, err18 := mbTrans17.WriteString(arg16)
		if err18 != nil {
			Usage()
			return
		}
		factory19 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt20 := factory19.GetProtocol(mbTrans17)
		argvalue0 := demo.NewArticle()
		err21 := argvalue0.Read(jsProt20)
		if err21 != nil {
			Usage()
			return
		}
		value0 := argvalue0
		fmt.Print(client.Put(value0))
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
