__global__ void cuda_hello (){
    printf("Hello From GPU!\n");
}

int main(){
    cuda_hello<<<1,1>>>();
    // CPU waits here for the GPU to finish
    cudaDeviceSynchronize();
    return 0;
}
