# Fehler-Zusammenfassung

- Vergleich `currentLen < longestLen<`
  mit der Aktualisierung danach macht longestLen kleiner.
  Das passt nicht zur Aufgabe, dass das längste Element gesucht wird.
- Die Bedingung `currentLen <= 3` ist genau falsch herum.
  Wenn die Länge kleiner als 3 ist, fängt das Element auch nicht mit
  "abc" an.
- Die Bedingung `val[1:]` scheint völlig willkürlich.
  Aus Tests abgeschrieben? Das wäre dann der Versuch einer
  hartcodierten Lösung.
