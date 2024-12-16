import binascii

s = "5B0B11114A120E0C08151D115D160E0705085B17070108071604190417"
b = bytearray(binascii.unhexlify(s))

iv = bytearray(b"samesamebutdifferentdifferent")

# /home/dlegezo/sec_file_generator.bpf.c
    # __builtin_memcpy(decrypted_msg, encrypted_msg, sizeof(encrypted_msg));
    #     data[k] ^= iv[28-k];
    # for (k = 0; k < 25; k++) {
    # bpf_printk("%s %d", decrypted_msg, counter);
    # counter++;
    # return XDP_PASS;

print(len(b), len(iv))
for k in range(29):
    b[k] ^= iv[28-k]

print(b)
