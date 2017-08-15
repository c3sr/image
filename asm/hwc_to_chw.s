	.section	__TEXT,__text,regular,pure_instructions
	.macosx_version_min 10, 12
	.intel_syntax noprefix
	.globl	_rgb_hwc_to_chw
	.p2align	4, 0x90
_rgb_hwc_to_chw:                        ## @rgb_hwc_to_chw
## BB#0:
	push	rbp
	mov	rbp, rsp
	push	r15
	push	r14
	push	r13
	push	r12
	push	rbx
	and	rsp, -8
                                        ## kill: %ECX<def> %ECX<kill> %RCX<def>
	test	edx, edx
	jle	LBB0_9
## BB#1:
	test	ecx, ecx
	jle	LBB0_9
## BB#2:
	mov	eax, ecx
	imul	eax, edx
	lea	ebx, [rax + rax]
	movsxd	r8, ebx
	movsxd	r15, eax
	lea	r9d, [rcx - 1]
	lea	r10, [r9 + 2*r9 + 3]
	mov	r11d, ecx
	and	r11d, 1
	xor	r14d, r14d
	.p2align	4, 0x90
LBB0_3:                                 ## =>This Loop Header: Depth=1
                                        ##     Child Loop BB0_7 Depth 2
	test	r11d, r11d
	mov	r12d, 0
	mov	rbx, rsi
	je	LBB0_5
## BB#4:                                ##   in Loop: Header=BB0_3 Depth=1
	mov	eax, dword ptr [rsi]
	mov	dword ptr [rdi], eax
	mov	eax, dword ptr [rsi + 4]
	mov	dword ptr [rdi + 4*r15], eax
	lea	rbx, [rsi + 12]
	mov	eax, dword ptr [rsi + 8]
	mov	dword ptr [rdi + 4*r8], eax
	mov	r12d, 1
LBB0_5:                                 ##   in Loop: Header=BB0_3 Depth=1
	test	r9d, r9d
	je	LBB0_8
## BB#6:                                ##   in Loop: Header=BB0_3 Depth=1
	mov	r13d, ecx
	sub	r13d, r12d
	.p2align	4, 0x90
LBB0_7:                                 ##   Parent Loop BB0_3 Depth=1
                                        ## =>  This Inner Loop Header: Depth=2
	mov	eax, dword ptr [rbx]
	mov	dword ptr [rdi], eax
	mov	eax, dword ptr [rbx + 4]
	mov	dword ptr [rdi + 4*r15], eax
	mov	eax, dword ptr [rbx + 8]
	mov	dword ptr [rdi + 4*r8], eax
	mov	eax, dword ptr [rbx + 12]
	mov	dword ptr [rdi], eax
	mov	eax, dword ptr [rbx + 16]
	mov	dword ptr [rdi + 4*r15], eax
	mov	eax, dword ptr [rbx + 20]
	mov	dword ptr [rdi + 4*r8], eax
	add	rbx, 24
	add	r13d, -2
	jne	LBB0_7
LBB0_8:                                 ##   in Loop: Header=BB0_3 Depth=1
	lea	rsi, [rsi + 4*r10]
	inc	r14d
	cmp	r14d, edx
	jne	LBB0_3
LBB0_9:
	xor	eax, eax
	lea	rsp, [rbp - 40]
	pop	rbx
	pop	r12
	pop	r13
	pop	r14
	pop	r15
	pop	rbp
	ret


.subsections_via_symbols
