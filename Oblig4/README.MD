# Obligatorisk oppgave 4 #

#### Gruppemedlemmer ####
* Jarl Andreassen 
* Vegard Saavesen Mathisen 
* Øyvind Markeng Hammer 
* Eivind Pedersen 


## Systemspesifikasjon ## 

Vår applikasjon tar for seg en daglig værmelding for tre bestemte byer, hvor byene vi har valgt er Oslo, Stavanger og Trondheim. Nytteverdien av denne applikasjonen er at bruker skal kunne få opplyst hvordan været er i det øyeblikket applikasjonen blir kjørt, og ut fra denne informasjonen vil bruker få en opplysning med informasjon om hvordan man burde kle seg, og hva man eventuelt bør ta med seg om man skal ut.

Vi får informasjon fra gjennomsnittsverdiene: Lufttrykket, luftfuktighet, vindstyrke og grader, samt direktebilder fra de respektive stedene. 

Lufttrykket måles i kilopascal, da det er måleenheten for lufttrykk. Luftfuktigheten returneres i prosent, fordi det er enklest å forstå. Vindstyrken er satt til m/s og gradene er satt til °C fordi det er det mest brukte i Norge. 
Sidene vi har hentet informasjon fra er:

* Openweathermap.org for lufttrykk, luftfuktighet og vindstyrke.
* Accuweather.com for gradestokken i celsius og klokkeslett. 
* Lookr.com for direktesendte bilder fra de respektive byene. 

Brukers nytteverdi, scenario: 

Scenario 1:

Jarl skal på Aker brygge og ta seg noen pils med sine kompiser. De siste dagene har det vært sol, men den siste uken har det vært litt mye vind. Jarl sjekker da vår applikasjon slik at han vet hvordan været er hva han eventuelt skulle trenge å kle på seg. Han kan også sjekke ut kameraer som sender live fra bryggen for å se om det er mange ute denne dagen. 

Scenario 2: 

Knut er i Forsvaret i Stavanger på Madla leir og skal på feltøvelse. Det er høst og veldig skiftende vær. Han er lei av å ta på seg for mye eller for lite klær, og sjekker derfor applikasjonen for å sjekke hvordan han skal kle seg og pakke feltsekken. Han må kle seg i regntøy og vindtett, siden applikasjonen melder nedbør og vind. 

  



systemarkitektur



# Installasjonsintruksjoner #
* For å kjøre programmet må alle go og html filene inn i samme mappe på golang. 
* Programmet kjøres ved kommandoen "go run vermelding.go" i terminalen.
