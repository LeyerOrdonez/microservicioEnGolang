package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//funcdel 3 metodo

type Pagina struct{
	Titulo string 
	Cuerpo []byte
}


func manejadorMostrarPagina(w http.ResponseWriter, r *http.Request){
	titulo := r.URL.Path
	pagina, _ := cargarPagina(titulo)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", pagina.Titulo[len("/pag/"):],pagina.Cuerpo)  //ruta de la pagina

    fmt.Println("pagina hecha")

}

func (p* Pagina)guardar() error{
	nombre := p.Titulo + ".txt"
	return ioutil.WriteFile("./pag/" + nombre,p.Cuerpo,0600)
}

func cargarPagina(titulo string)(*Pagina, error){
	nombre_archivo := titulo + ".txt"
	fmt.Println("el cliente ha pedido"+ nombre_archivo)
	cuerpo, err := ioutil.ReadFile("."+nombre_archivo)

	if(err != nil){
		return nil,err
	}

	return &Pagina {Titulo: titulo, Cuerpo: cuerpo},nil
}



//funcion aparte que lee archivos 
func leerArchivo(text string) string {
	datosComoBytes, err := ioutil.ReadFile(text)
	if err != nil{
		log.Fatal(err)
	}
	datosComoString := string(datosComoBytes)
	return datosComoString
}

//carga paginas del disco (no funcional)
func cargarPaginaDisk(titulo string)(*Pagina , error){
	nombre_archivo := titulo +".html"
	cuerpo,err := ioutil.ReadFile(nombre_archivo)
	if(err != nil){
		return nil,err
	}

	return &Pagina {Titulo: titulo, Cuerpo: cuerpo},nil

}




func main() {


       //metodo no comprobado
    // http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
    //     fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    // })

    //  http.HandleFunc("/MicroservicioLocal",  func(rw http.ResponseWriter, r *http.Request){ 
	// 	html := blackfriday.MarkdownCommon([]byte (r.FormValue("cuerpo")))
	//     rw.Write(html)
	//  })

    // http.Handle("/", http.FileServer(http.Dir("/publico/index.html")))

       //metodo comprobado y funcional, hata cierta parte
	// http.HandleFunc("/MicroservicioLocal",  func(rw http.ResponseWriter, r *http.Request){ 
	// 	fmt.Fprintf(rw,"<!DOCTYPE html><html lang=\"es\"><head><meta charset=\"utf-8\">	<title>Generador de HTML</title>	<link href=\"/css/bootstrap.min.css\" rel=\"stylesheet\"> </head> <body><div class=\"container\"><div class=\"page-title\">	<h1>Editor de texto en formato markdown</h1><p class=\"lead\">Go genera HTML a partir de markdown</p><hr /> </div> <form action=\"/MicroservicioLocal\" method=\"POST\"><div class=\"form-group\">  <textarea class=\"form-control\" name=\"cuerpo\" cols=\"30\" rows=\"10\"></textarea></div><div class=\"form-group\"><input type=\"submit\" class=\"btn btn-primary pull-right\" /></div> </form></div><script src=\"/js/bootstrap.min.js\"></script>  </body></html>")
	    
	// })


    //  fmt.Println("servidor en funcionamiento!")
    // log.Fatal(http.ListenAndServe(":8081", nil))


	    //3 metodo 
	pag1 := &Pagina{Titulo:"Autoservice Fourier", Cuerpo:[]byte( leerArchivo("./pag/index.html") )}
	pag1.guardar()
	http.HandleFunc("/pag/", manejadorMostrarPagina)
	fmt.Println("servidor en funcionamiento!")
	log.Fatal(http.ListenAndServe(":8081", nil))

    
}
