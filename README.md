# Scraper-Quotes
Proyecto de aprendizaje enfocado en los fundamentos de Go y Web Scraping. El scraper recorre de forma recursiva múltiples páginas, gestiona la persistencia de datos en JSON y aplica técnicas de evasión de detección.

Este es mi primer proyecto en ***Go***, lo diseñe para aprender las bases del Web Scraping  y para aprender a utilizar el lenguaje de ***Go***

Utilice las siguientes tecnologias y utilidades para armar mi proyecto:
**Go** 
**Colly**
**Json**

El Programa tiene las siguientes caracteristicas:
Navegacion Recursiva : Avanza de pagina hasta que llegue a la ultima
Proteccion Anti-Bot: Esta caracterista la implemente solo para aprender sobre los metodos de proteccion, la pagina de Quotes no tiene muchas protecciones contra el scrapeo (si no es que ninguna)
Delay : Tiene un delay incorporado para evitar que el programa recorra todo en milisegundos y levante sospechas
Exportacion de datos: Guardo los datos en un archivo Json
