package main

//
//func Logger(inner rootHandler, name string) rootHandler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		start := time.Now()
//
//		inner.ServeHTTP(w, r)
//
//		log.Printf(
//			"%s\t%s\t%s\t%s",
//			r.Method,
//			r.RequestURI,
//			name,
//			time.Since(start),
//		)
//	})
//}