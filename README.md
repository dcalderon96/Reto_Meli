# Challenge Backend - Meli
1. [Descripción](#Descripción-)
2. [Configuración del proyecto](#configuraci%C3%B3n-del-proyecto-%EF%B8%8F)
3. [Utilidades](#Utilidades-)
4. [Ejecución del proyecto](#ejecuci%C3%B3n-del-proyecto-%EF%B8%8F)
5. [Tecnologías](#Tecnologías-)
6. [Autor](#Autor-)

## Descripción 📖
Challenge Backend realizado para implementar un programa y un API que permita determinar el posicionamiento de una nave en el espacio exterior, asi mismo permitira conocer el mensaje de auxilio emitido a partir de tres satelites que reciben la transmisión y conocen la distancia a la cual se encuentra la nave emisora.

## Configuración del proyecto ⚙️
Git:

    git clone https://github.com/dcalderon96/Reto_Meli.git
    cd Reto_Meli

o descarga el archivo [zip](https://github.com/dcalderon96/Reto_Meli/archive/refs/heads/master.zip).

## Utilidades 🔍
En el repositorio se encuentra el archivo **Reto_Meli.postman_collection**, el cual tiene como finalidad ayudar a realizar las pruebas de consumo del API.

## Ejecución del proyecto ▶️
1. Para la ejecución de la aplicación Shell se debe ejecutar el archivo *main.exe* ubicado en la siguiente ruta *Reto_Meli/Program/bin*
2. Para usar el API de froma local se debe ejecutar el siguiente archivo: *main.exe* ubicado en la siguiente ruta *Reto_Meli/API/bin*, el cual iniciara un servidor local ubicado en *(localhost:8080)*

3. Abrir la colección de Postman y verificar el modo de ejecución que usaremos online o local (se debe haber realizado el paso 2 previamente), una vez seleccionado el modo de ejecución se debe importar la siguiente coleccion en POSTMAN: *Reto_Meli/Reto_Meli.postman_collection*

Nota:
```
Verificar la url del API en la barra de cada método, en caso de que ejecutemos el API localmente cambiar la url por: *localhost:8080"
``` 

## Tecnologías 💻
- Go

## Autor 👨
Daniel Felipe Calderón <d.calderon96.dc@gmail.com>