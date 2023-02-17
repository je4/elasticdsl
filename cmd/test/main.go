package main

import (
	"crypto/tls"
	"emperror.dev/emperror"
	"emperror.dev/errors"
	"flag"
	"github.com/je4/elasticdsl/v2/pkg/elastic"
	"github.com/je4/elasticdsl/v2/pkg/elastic7"
	"github.com/je4/elasticdsl/v2/pkg/elastic8"
	lm "github.com/je4/utils/v2/pkg/logger"
	"net/http"
)

const logFormat = `%{time:2006-01-02T15:04:05.000} %{module}::%{shortfunc} [%{shortfile}] > %{level:.5s} - %{message}`

var _elasticendpoint = flag.String("elastic", "", "endpoint for elastic search")
var _elasticAPIKey = flag.String("elasticapikey", "", "apikey for elasticsearch version 8 client")
var _elastic7 = flag.Bool("elastic7", false, "use elastic search api v7")

func GetErrorStacktrace(err error) errors.StackTrace {
	type stackTracer interface {
		StackTrace() errors.StackTrace
	}

	var stack errors.StackTrace

	errors.UnwrapEach(err, func(err error) bool {
		e := emperror.ExposeStackTrace(err)
		st, ok := e.(stackTracer)
		if !ok {
			return true
		}

		stack = st.StackTrace()
		return true
	})

	if len(stack) > 2 {
		stack = stack[:len(stack)-2]
	}
	return stack
	// fmt.Printf("%+v", st[0:2]) // top two frames
}

func main() {
	flag.Parse()

	// create logger instance
	logger, lf := lm.CreateLogger("Client DSL Test", "", nil, "DEBUG", logFormat)
	defer lf.Close()

	var eClient elastic.Client
	var err error
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	if *_elastic7 {
		eClient, err = elastic7.NewClient(*_elasticendpoint, "", *_elasticAPIKey, logger)
	} else {
		eClient, err = elastic8.NewClient(*_elasticendpoint, "", *_elasticAPIKey, logger)
		if err != nil {
			logger.Errorf("cannot connect to eClient endpoint '%s': %v", *_elasticendpoint, err)
			logger.Debugf("%v%+v", err, GetErrorStacktrace(err))
			return
		}
	}
	serverInfo, err := eClient.Info()
	if err != nil {
		logger.Errorf("cannot get info from '%s': %v", *_elasticendpoint, err)
		logger.Debugf("%v%+v", err, GetErrorStacktrace(err))
		return
	}
	eClient.Search(&elastic.SearchConfig{
		Index:          "alma-je-test",
		Fields:         nil,
		QStr:           "",
		FiltersFields:  nil,
		Facets:         nil,
		Groups:         nil,
		ContentVisible: false,
		Start:          0,
		Rows:           0,
		IsAdmin:        false,
	})
	logger.Infof("Server: %v", serverInfo)
}
