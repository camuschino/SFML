# Laberth - SFML

Juego de laberintos hecho en C++ con SFML.

 SFML se descarga [aqui](https://www.sfml-dev.org/download/sfml/2.5.1). Se lo instala para que los modulos se puedan usar.

se crea un archivo de texto (con el nombre main, por ejemplo) y se le da la extension .cpp para programar en c++ y demas!

```
nano main.cpp
```

Si se ha instalado SFML directo de manejador de paquetes se puede compilar y crear Exe haciendo lo siguiente

Se compila el .cpp, y se genera un archivo con el mismo nombre con extension .o que es el codigo maquina.

```
g++ -c main.cpp
```

Ahora se genera un exe con el .o y los modulos de SFML que se usaron en el programa.
```
g++ main.o -o NOMBRE_DE_PROGRAMA -lsfml-graphics -lsfml-window -lsfml-system
```
Si se ha descargo el paquede desde la pagina y se esta usando el paquete de manera independiente se hace de la siguiente forma.

Se compila el .cpp, y se genera un archivo con el mismo nombre con extension .o que es el codigo maquina y se agrega el lugar donde esta el PATH de SFML.

```
g++ -c main.cpp -I <sfml-install-path>/include
```

Ahora se genera un exe con el .o y los modulos de SFML que se usaron en el programa. Y se deben indicar donde buscar las librerias que se han usado.

```
g++ main.o -o NOMBRE_DE_PROGRAMA -L <sfml-install-path>/lib -lsfml-graphics -lsfml-window -lsfml-system
```
