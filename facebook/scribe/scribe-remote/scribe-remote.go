// Autogenerated by Thrift Compiler (1.0.0-dev)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
	"facebook/scribe"
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
	fmt.Fprintln(os.Stderr, "  ResultCode Log( messages)")
	fmt.Fprintln(os.Stderr, "  string getName()")
	fmt.Fprintln(os.Stderr, "  string getVersion()")
	fmt.Fprintln(os.Stderr, "  fb_status getStatus()")
	fmt.Fprintln(os.Stderr, "  string getStatusDetails()")
	fmt.Fprintln(os.Stderr, "   getCounters()")
	fmt.Fprintln(os.Stderr, "  i64 getCounter(string key)")
	fmt.Fprintln(os.Stderr, "  void setOption(string key, string value)")
	fmt.Fprintln(os.Stderr, "  string getOption(string key)")
	fmt.Fprintln(os.Stderr, "   getOptions()")
	fmt.Fprintln(os.Stderr, "  string getCpuProfile(i32 profileDurationInSec)")
	fmt.Fprintln(os.Stderr, "  i64 aliveSince()")
	fmt.Fprintln(os.Stderr, "  void reinitialize()")
	fmt.Fprintln(os.Stderr, "  void shutdown()")
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
	client := scribe.NewScribeClientFactory(trans, protocolFactory)
	if err := trans.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
		os.Exit(1)
	}

	switch cmd {
	case "Log":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "Log requires 1 args")
			flag.Usage()
		}
		arg6 := flag.Arg(1)
		mbTrans7 := thrift.NewTMemoryBufferLen(len(arg6))
		defer mbTrans7.Close()
		_, err8 := mbTrans7.WriteString(arg6)
		if err8 != nil {
			Usage()
			return
		}
		factory9 := thrift.NewTSimpleJSONProtocolFactory()
		jsProt10 := factory9.GetProtocol(mbTrans7)
		containerStruct0 := scribe.NewLogArgs()
		err11 := containerStruct0.ReadField1(jsProt10)
		if err11 != nil {
			Usage()
			return
		}
		argvalue0 := containerStruct0.Messages
		value0 := argvalue0
		fmt.Print(client.Log(value0))
		fmt.Print("\n")
		break
	case "getName":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetName requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetName())
		fmt.Print("\n")
		break
	case "getVersion":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetVersion requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetVersion())
		fmt.Print("\n")
		break
	case "getStatus":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetStatus requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetStatus())
		fmt.Print("\n")
		break
	case "getStatusDetails":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetStatusDetails requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetStatusDetails())
		fmt.Print("\n")
		break
	case "getCounters":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetCounters requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetCounters())
		fmt.Print("\n")
		break
	case "getCounter":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCounter requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetCounter(value0))
		fmt.Print("\n")
		break
	case "setOption":
		if flag.NArg()-1 != 2 {
			fmt.Fprintln(os.Stderr, "SetOption requires 2 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		argvalue1 := flag.Arg(2)
		value1 := argvalue1
		fmt.Print(client.SetOption(value0, value1))
		fmt.Print("\n")
		break
	case "getOption":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetOption requires 1 args")
			flag.Usage()
		}
		argvalue0 := flag.Arg(1)
		value0 := argvalue0
		fmt.Print(client.GetOption(value0))
		fmt.Print("\n")
		break
	case "getOptions":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "GetOptions requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.GetOptions())
		fmt.Print("\n")
		break
	case "getCpuProfile":
		if flag.NArg()-1 != 1 {
			fmt.Fprintln(os.Stderr, "GetCpuProfile requires 1 args")
			flag.Usage()
		}
		tmp0, err16 := (strconv.Atoi(flag.Arg(1)))
		if err16 != nil {
			Usage()
			return
		}
		argvalue0 := int32(tmp0)
		value0 := argvalue0
		fmt.Print(client.GetCpuProfile(value0))
		fmt.Print("\n")
		break
	case "aliveSince":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "AliveSince requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.AliveSince())
		fmt.Print("\n")
		break
	case "reinitialize":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Reinitialize requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Reinitialize())
		fmt.Print("\n")
		break
	case "shutdown":
		if flag.NArg()-1 != 0 {
			fmt.Fprintln(os.Stderr, "Shutdown requires 0 args")
			flag.Usage()
		}
		fmt.Print(client.Shutdown())
		fmt.Print("\n")
		break
	case "":
		Usage()
		break
	default:
		fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
	}
}
