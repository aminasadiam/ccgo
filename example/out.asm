global _start
_start:
    call main
    mov rdi, rax
    mov rax, 60
    syscall

main:
    push rbp
    mov rbp, rsp
    mov rax, 5
    mov rbx, rax
    mov rax, 3
    add rax, rbx
    mov rsp, rbp
    pop rbp
    ret
