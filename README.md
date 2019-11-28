# NoticiasArg-go
Simple go command-line project that retrieves the most relevant news from Argentina and shows them to you on screen.
It uses the Google News API through [newsapi](http://newsapi.org).

### Instructions
#### Requirements
1. Go installed
2. API key from [newsapi](http://newsapi.org). 
    Inside the project folder, create a file called ".env" and inside it, put
    ```
    NEWSAPI_KEY=your-key-here
    ```
3. Go to the project folder and run!

    ```
    go run noticias.go
    ```

    The output should look something like this
    ```
    Titulo: Fuerte apoyo de los industriales a Alberto Fernández, aunque las expectativas sobre la economía son moderadas 
    Desc: <ol><li>Fuerte apoyo de los industriales a Alberto Fernández, aunque las expectativas sobre la economía son moderadas  infobae
    </li><li>Las 15 definiciones económicas de Alberto Fernández  ámbito.com
    </li><li>Empresarios valoran el mensaje de Alberto pero p…
    Link: https://www.infobae.com/economia/finanzas-y-negocios/2019/11/28/fuerte-apoyo-de-los-industriales-a-alberto-fernandez-aunque-las-expectativas-sobre-la-economia-son-moderadas/ 
    
    Titulo: Incidentes en la Legislatura porteña por un proyecto que regula la actividad de profesionales de la salud en hospitales públicos 
    Desc: Médicos, residentes, enfermeros y concurrentes de salud se movilizaron para expresar su rechazo a un proyecto
    Link: https://www.infobae.com/politica/2019/11/28/incidentes-en-la-legislatura-portena-por-un-proyecto-que-regula-la-actividad-de-profesionales-de-la-salud-en-hospitales-publicos/ 
    
    Titulo: Por 30 mil votos, Uruguay da un giro a la derecha con Luis Lacalle Pou | Opinión 
    Desc: <ol><li>Por 30 mil votos, Uruguay da un giro a la derecha con Luis Lacalle Pou | Opinión  Página 12
    </li><li>Uruguay: el escrutinio definitivo consagró a Luis Lacalle Pou como nuevo Presidente  infobae
    </li><li>Uruguay: ¿quién es Luis Lacalle Pou y qué espe…
    Link: https://www.pagina12.com.ar/233583-por-30-mil-votos-uruguay-da-un-giro-a-la-derecha-con-luis-la 
    ```