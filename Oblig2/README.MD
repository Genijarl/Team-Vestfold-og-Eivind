# Team Vestfold og Farsund

Jarl Andreassen: Jarla17@uia.no

Vegard Mathisen: Vegasm17@uia.no

Eivind Pedersen: Eivip17@uia.no

Øyvind Hammer Markeng: Oyvinm15@uia.no



# OBLIG2

Oppgave 1
a) Skriv et Golang program med navn fileinfo.go​,

Programmet skal returnere detaljert informasjon om en fil i følgende format:

Information about file <filnavn>:

Size: X in bytes, KB, MB and GB

 Is/Is not a directory

 Is/Is not a regular file

 Has Unix permission bits: -rwxrwxrwx

 Is/Is not append only

 Is/Is not a device file

 Is/Is not a Unix character device

 Is/Is not a Unix block device

 Is/Is not a symbolic link
b) Lag en kjørbart program (build) av fileinfo.go som kan utføres fra kommandolinje. Programmet skal ta et argument, som skal være et filnavn.

Kommandoet for å utføre programmet i terminal skal se slik ut: fileinfo

Stien til filnavnet kan være enten relativ eller absolutt.

Link til oppgave 1: 
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave1/fileinfo.go

Oppgave 2
a) Skriv programmer med navn filecount.go som skal lese en tekst-fil, skrive ut totalt antall linjer og de fem “runes” (bokstaver/tegn), som forekommer mest i filen samt hvor mange ganger de telles i den medfølgende filen text.txt. Som forrige oppgave skal filen bygges til et kjørbart program og kunne utføres i terminal med kommandoen: filecount

Resultatet skal være i følgende format:

Information about <filnavn>:

Number of lines in file:  X

Most common runes:

Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X
Rune: X , Counts: X

Link til oppgave 2: 
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave2/filecount.go

Oppgave 3
a) Skriv et program med navn addup.go bestående av to go funksjoner og kommunisere med hverandre gjennom channels

Funksjon A:

Lag en go funksjon som leser inn to tall fra terminal (stdin) og sender disse to verdiene til en annen go funksjon.

Funksjon B:

Mottakende go funksjon skal legge sammen de to tallene og sende resultatet tilbake til funksjon A som skriver ut resultatet til terminal (stdout).

b) Skriv to programmer (addtofile.go og sumfromfile.go) som kommuniserer med hverandre gjennom en fil.

addtofile.go: Leser inn to tall og skriver disse til en fil.

sumfromfile.go: Leser tall fra fil laget av program A og adderer de to tallene. Summen skrives tilbake til samme fil. Program A leser resultatet fra fil og skriver ut summen til stdout.

c) Beskriv og implementer feilhåndtering på all I/O i oppgave a) og b). Se https://blog.golang.org/error-handling-and-go (Lenker til en ekstern side.)Lenker til en ekstern side. for mer informasjon om feilhåndtering i Go.

d) Implementer håndtering av SIGINT i programmene i a) og b); programmene skal skrive ut en avslutningsmelding dersom de mottar SIGINT før de fullfører naturlig.

e) Bygg filene som i oppgave 1 og 2 til kjørbare filer på ditt operativsystem og legg dem ved i /bin mappen.

Link til oppgave 3a: 
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/addup.go

Linker til oppgave 3b og 3d: 
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/addtofile.go
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/sumfromfile.go

Link til oppgave 3c:
https://github.com/Genijarl/Team-Vestfold-og-Farsund/blob/master/Oblig2/src/oppgave3/Feilh%C3%A5ndtering

Link til oppgave 3e:
https://github.com/Genijarl/Team-Vestfold-og-Farsund/tree/master/Oblig2/src/bin
