# pf-neame: ¡El Generador de Pesadillas para QA!

¡Bienvenidos al infierno burocrático! Esta "utilidad" CLI es un chiste cruel para el equipo de QA, diseñado para que se "pf-nemee" (léase: "fuck me pf-xx document") con archivos .dll y .exe. Básicamente, es un "pf-neame este" en español, porque ¿quién no ama diligenciar documentos cuando podría estar probando bugs reales?

## ¿Qué hace esta maravilla?
Toma la información de versión de tus archivos (comentarios, fileversion y demás) y automáticamente diligencia los sagrados documentos PF-26 y PF-30. ¡Porque nada dice "productividad" como automatizar el papeleo que nadie quiere hacer!

## Instalación
Clona este repo, instala Go (si no lo tienes, ¡qué pena!), y corre:
```
go build -o pf-neame main.go
```

## Uso
Ejecuta el comando mágico:
```
./pf-neame <ruta-al-archivo.dll-o-exe>
```
Y voilà! Tus documentos PF-26 y PF-30 se llenarán solos. ¡O no! (Spoiler: probablemente sí, pero con un toque de caos para mantener el humor negro).

## Ejemplos
- `./pf-neame miApp.exe` → Genera PF-26 con versión 1.0.0 y comentarios "Esta app es un chiste".
- `./pf-neame --help` → Muestra ayuda, porque incluso los chistes necesitan instrucciones.

## Contribuciones
¿Quieres agregar más sufrimiento? ¡Pull requests welcome! Pero recuerda: esto es un joke, no lo tomes en serio. Si encuentras bugs, es porque el QA no lo probó bien. 😈

## Licencia
MIT, porque incluso los chistes necesitan licencias. ¡Diviértete pf-nemeando!
