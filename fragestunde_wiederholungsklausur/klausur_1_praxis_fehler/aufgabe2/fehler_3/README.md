# Fehler-Zusammenfassung

- `result` wird als leere Liste initialisiert und dann läuft
  die Schleife darüber. Diese Schleife läuft tatsächlich nie,
  `result` bleibt also immer leer.
- `firstposition` und `lastposition` kommen nicht unbedingt in
  der Liste vor (siehe zweiter Testfall).
  Der Ansatz war offenbar aus einer Probeklausur abkopiert,
  funktioniert hier aber nicht.
