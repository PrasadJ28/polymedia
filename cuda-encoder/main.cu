#include <stdio.h>

__global__ void cuda_hello (){
    printf("Hello From GPU!\n");
}

int main(){
    cuda_hello<<<1,1>>>();
    cudaDeviceSynchronize();
    int device;
    cudaGetDeviceCount(&device);
    cudaDeviceProp props;
    cudaGetDeviceProperties(&props, 0);
    printf("Device Name: %s\n", props.name);
    float size = props.totalGlobalMem/(1024*1024*1024);
    printf("Device Size%2f GB\n",size);
    printf("Max Threads per Block %d\n", props.maxThreadsPerBlock);
    return 0;
}
