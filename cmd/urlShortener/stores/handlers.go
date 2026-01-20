package stores

import (
	"fmt"
	"net/http"
)



func Add( w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html")
	url:=r.FormValue("url")
	// fmt.Print("url: ", url)
	if url==""{
		fmt.Fprint(w,AddForm)
		return
	}
	key:=urlStore.Put(url)
	fmt.Print("key: ", key)
	fmt.Fprintf(w , "%s",key)
}

const AddForm =` 
<html>
<body>
<form method="POST" action="/add">
URL:<input type="text" name="url">
<input type="submit" value="Add">
</form>
</body>
</html>`


// url : /abc , the key would be - abc
func Redirect(w http.ResponseWriter, r *http.Request){
	 
	key:= r.URL.Path[1:]
	url := urlStore.Get(key)
	if url==""{
		http.NotFound(w, r)
		return
	}
	http.Redirect(w, r, url, http.StatusFound)
}