# Obligatorisk oppgave 3 #

#### Gruppemedlemmer ####
* Jarl Andreassen *
* Vegard Saavesen Mathisen *
* Øyvind Markeng Hammer *
* Eivind Pedersen *


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
