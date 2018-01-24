# godemo
godemo
This is practise of golang to profile the "hello world" http service, and compare the difference between net/http and valyala/fasthttp.

Current conclusion is.

with client http session keep alive.
valyala/fasthttp is very much better than net/http.

without client http session keep alive.
valyala/fasthttp get high error rate since the port recycle speed cannot catch up with the client session check in. 
and in this condition net/http wins with it stableness.
