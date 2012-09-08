Web-palvelinohjelmointi s2012
=============================

tl;dr:
( Go tulee löytyä koneelta etukäteen, http://golang.org/doc/install )

Testien ajo kussakin kansiossa (suorittaa *_test.go tiedostoissa sijaitsevat testit)
    go test

Ohjelmat jotka ovat yhdessä tiedostossa, esim server.go kokonaisuudessaan voi suorittaa ilman säätöä komennolla
    go run server.go

Ohjelman voi myös kääntää (ja sen jälkeen suorittaa tiedoston nimellä toki)
    go build

Ps. Kaikki serverit taitavat tällä hetkellä käyttää oletuksena porttia 8080 ellei toisin mainita.
Ps2. Serverien kuuntelemat polut voi lukea sorsasta. Pyrin mainitsemaan ne tiedoston alussa, handler kertoo toki varmasti.
