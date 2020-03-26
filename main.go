  
  func main() {
    	router := mux.NewRouter()
    	router.HandleFunc("/gomk/fixall/{server}", handler).Methods("GET")
    	headers := handlers.AllowedHeaders([]string
        {"X-Requested-With", "Content-Type", "Authorization", 
        "access-control-allow-headers", "access-control-allow-methods",
         "access-control-allow-origin"})
    	methods := handlers.AllowedMethods([]string
        {"GET", "POST", "PUT", "DELETE", "OPTIONS"})
    	origins := handlers.AllowedOrigins([]string{"*"})
    	fmt.Println("Server is running at port 8000")
    	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(headers, 
        methods, origins)(router)))
    }
    func handler(w http.ResponseWriter, r *http.Request) {
    	params := mux.Vars(r)
    	server := params["server"]`
    resp, err := http.Get(`https://test.com/api/test&request={"hostname":"` 
    + params["server"] + `","mode":"test"}`)
    	if err != nil {
    		log.Fatalln(err)
    		return
    	}
    	defer resp.Body.Close()
    	body, err := ioutil.ReadAll(resp.Body)
    	if err != nil {
    		log.Fatalln(err)
    	}
    	json.NewEncoder(w).Encode(string(body))
    	fmt.Println(string(body))
    }
