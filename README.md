# Team-Vestfold-og-Farsund

Jarl Andreassen   Jarla17@uia.no

Vegard Mathisen   Vegasm17@uia.no

Eivind Pedersen   Eivip17@uia.no

Øyvind Hammer     Oyvinm15@uia.no

# Obligatorisk oppgave 1

# Oppgave 1. 
Binære tall, Hexadesimaltall og Desimaltall

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/1AB

# Oppgave 2.
Forstå algoritmer og utføre “benchmark”-tester på koden



https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/2ABC

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/sorting.go

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/sorting_test.go

<img src="https://user-images.githubusercontent.com/35763714/36151603-231996f6-10c8-11e8-8cc2-9e7f53c86ebd.png" width="90%"></img> 

# Oppgave 3
Forstå prosessadministrajon på et platform

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/3

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/evigl%C3%B8kke.go



# Oppgave 4
Typografiske symboler

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/4ABC

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/ascii.go

https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig1/iso_test.go

# Obligatorisk oppgave 2

# Oppgave 1
Go build fileinfo.go og bruk fileinfo som argument med påfølgende filnavn. eks: fileinfo text.txt
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/fileinfo.go

# Oppgave 2
Go build filecount.go og bruk filecount som argument med påfølgende filnavn. eks: filecount text.txt
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/filecount.go

# Oppgave 3
a) + d)
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/addup.go
b) + d)
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/addtofile.go
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/sumfromfile.go

c) Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b). Se https://blog.golang.org/error-handling-and-go (Lenker til en ekstern side.)Lenker til en ekstern side. for mer informasjon om feilhåndtering i Go.

d) Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig. 
 https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/addtofile.go 
 https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/sumfromfile.go

e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.


# Obligatorisk oppgave 3 #

#### Gruppemedlemmer ####
* Jarl Andreassen 
* Vegard Saavesen Mathisen 
* Øyvind Markeng Hammer 
* Eivind Pedersen 


## Oppgave 1 ##

*a)* Sett opp en webserver som lytter til port 8080.

*b)* Når klienter aksesserer webserveren på path lik "/" skal klienten få tilbake strengen "Hello, client".
Strengen skal vises i nettleseren.

https://github.com/Genijarl/Team-Vestfold-og-Farsund/tree/master/Oblig3/src/Oppgave%201%20og%202


## Oppgave 2 API kall og behandling av JSON og HTML templates ##

*a)* Presenter data på webserveren slik at den er leselig for mennesker(Ren tekst, f.eks. "Sted: Jorenhholem, dato: 13.04.2018", 
ikke i curly brackets.) når klienter aksesserer stiene /1, /2, /3, /4, /5.

Dataen som skal presenteres skal hentes fra fem ulike APIer, hvor alle returnerer data i JSON format. Dere velger selv hvilke datasett 
dere benytter. Det er denne dataen som skal presenteres på de ulike stiene på webserveren.

*b)* Samtlige stier i oppgave 2 skal rendres til klienter ved bruk av Go templates.

https://github.com/Genijarl/Team-Vestfold-og-Farsund/tree/master/Oblig3/src/Oppgave%201%20og%202

 

## Oppgave 3 UDP/TCP server og Internett standarder ##

Implementer et serverprogram i henhold til RFC 865. Serveren skal svare både på UDP og TCP.

https://github.com/Genijarl/Team-Vestfold-og-Farsund/tree/master/Oblig3/src/Oppgave%203


FYI: Når serveren lytter etter tilkoblinger på port 17, så fungerer det ved hjelp av kommandoen "telnet localhost 17". Dette gjelder TCP,
men når det kommer til UDP så må NETCAT lastes ned. Dette er fordi telnet kun fungerer til TCP, mens NETCAT fungerer til begge deler. Vi 
pratet med Martin, han sa dette var godt nok.


