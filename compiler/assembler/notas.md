## Primer paso Instalar qemu

```bash
sudo apt install qemu-user-static
```

## Segundo paso  

### Opción 1 (Utilizando Makefile):

Correr el archivo Makefile con el comando
```bash
make program
```
### Opción 2 (Utilizando build.sh):

Correr el archivo Makefile con el comando
```bash
chmod +x build.sh
```
```bash
./build.sh program.s
```

## Tercer Paso

### Ejecutar qemu directamente y correr el ejecutable

Correr el comando:
```bash
qemu-aarch64 ./program
```

---------------------------- Opcional --------------------------------------------------

### Ejecutar qemu en una terminal y utilizar gdb

Correr el comando:
```bash
qemu-aarch64 -g 1234 ./program
```


## Cuarto paso (Abrir en otra terminal para ejecutar gdb)

```bash
gdb-multiarch -q --nh \
  -ex 'set architecture aarch64' \
  -ex 'file hello-world' \
  -ex 'target remote localhost:1234' \
  -ex 'layout split' \
  -ex 'layout regs'
```

## Comandos para gdb

- step = Esto es para programas en C/C++, o para puntos de interrupción (breakpoints). Ejecuta el programa y lo detiene en la siguiente línea del código fuente en C o en el siguiente punto de interrupción (si se ha establecido uno).

- stepi = Ejecuta una sola instrucción (solo para la vista en ensamblador).

- nexti = Ejecuta la siguiente instrucción inmediatamente inferior; todas las bifurcaciones (branches) se tratan como instrucciones nop (sin operación).

- break [nombre de función] = Establece un punto de interrupción en una función.

- next = Permite que el programa se ejecute hasta el siguiente punto de interrupción.

- continue = Continúa la ejecución del programa desde donde está detenido; se volverá a detener si se alcanza otro punto de interrupción.

- quit = Salir de GDB.

- delete = Eliminar todos los puntos de interrupción.

- info vector = Muestra los registros de punto flotante (FPRs) primero como datos vectoriales y luego como datos de punto flotante.

- layout next = Cambia a una vista diferente (de C/C++ a ensamblador, por ejemplo).

- layout prev = Regresa a la vista anterior.