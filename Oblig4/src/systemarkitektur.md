## Systemarkitektur ##

Webserveren vår henter inn oppdaterte data fra forskjellige værmeldings API’er hos openweathermap. 
Disse går fra henholdsvis Oslo, Trondheim og Stavanger. Når brukeren bruker tjenesten vil brukeren se resultatet av dataen programmet
henter fra disse API’ene i nettleseren slik at det blir menneskelig å lese. 
Dette fungerer slik:

![image](https://user-images.githubusercontent.com/35763762/39866250-245c0f64-5450-11e8-8964-e6d2cb04b93d.png)

Dette viser at for hver gang vi henter en API vil dette vises i serveren vår i form av reelle tall og bynavn i et redesignet
miljø tilpasset en nettside. Nettsiden henter widgets for å gi en bedre interface av systemet. 
Når programmet kjøres starter det nederst i hiearkiet ved å hente data, så hente widgets og deretter kommer resultatet
opp i web serveren vår. Dette gjøres ved å bruke en lokal server (localhost:8080). Webserveren vår henter data fra de forskjellige
API'ene og sender de tilbake til brukeren etter tallene og widgets er hentet fra sine respektive domener.
