# hello.go

Analyze the following assembly language step by step.

> ````assembly
> 0000000000401000 <internal/cpu.Initialize>:
> 
> // Initialize examines the processor and sets the relevant variables above.
> // This is called by the runtime package early in program initialization,
> // before normal init functions are run. env is set by runtime if the OS supports
> // cpu feature options in GODEBUG.
> 
> func Initialize(env string) {
> ````

- `GODEBUG` is an environment variable for the Go runtime. It is used to debug Go programs and set runtime behavior. 
  
  ```bash
  # Print detailed information for each garbage collection
  $ GODEBUG=gctrace=1 go run main.go 
  ```
  
  
  
- From the comments, we can understand that `the Initialize function is called earlier` by the runtime package during program initialization, even earlier than normal init functions.

- The name of this function is internal/cpu.Initialize, which initializes `CPU-related content`. 

> ```assembly
> 401000: 49 3b 66 10 cmp 0x10(%r14),%rsp
> 401004:	76 38       jbe 40103e <internal/cpu.Initialize+0x3e>
> ```

- `Compare` the values of `rsp` and `0x10(%r14)`. If less than or equal to, execute the subroutine at 40103e.

- Before the Initialize function, there `may be other functions called earlier`. When executing, these earlier called functions `may have already placed some values at the 0x10 offset` associated with register r14.In the addressing mode of assembly language, % represents a register.

- `jbe` is an abbreviation for `jump if below or equal`.

> ```assembly
> 401006: 48 83 ec 18    sub $0x18,%rsp
> 40100a: 48 89 6c 24 10 mov %rbp,0x10(%rsp)
> 40100f: 48 8d 6c 24 10 lea 0x10(%rsp),%rbp
> ```

- Subtract 0x18 from rsp to allocate space for local variables on the stack.
- Store the value of rbp at 0x10(%rsp) to save the old base pointer value.
- Set 0x10(%rsp) to rbp to set a new stack base pointer.

>```assembly
>401014: 48 89 44 24 20 mov %rax,0x20(%rsp)
>401019: 48 89 5c 24 28 mov %rbx,0x28(%rsp)
>```

- The values in registers rax and rbx will be changed in subsequent code, but these two values are still needed at some point later.


> ```assembly
> doinit()
> ```

- Find the memory location of the doinit() subroutine.

> ```assembly
> 40101e:	66 90 xchg %ax,%ax
> ```

- The xchg %ax,%ax instruction exchanges the values of the ax register and the ax register.
- Since the ax register always exchanges its own value, this instruction actually has no effect and is a `no-op instruction`. 

> ````assembly
>   401020:	e8 9b 05 00 00       	callq  4015c0 <internal/cpu.doinit>
> ````

- Call the doinit() subroutine.

