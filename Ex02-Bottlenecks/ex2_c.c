#include <pthread.h>
#include <stdio.h>
pthread_mutex_t mutex1 = PTHREAD_MUTEX_INITIALIZER;
int i=0;
// Note the return type: void*
void* add(){
     int k=0;
     for (k; k<1000000; k++){
	  pthread_mutex_lock( &mutex1);
	  i++;
	  pthread_mutex_unlock( &mutex1);
     }
}

void* sub(){
     int j=0;
     for (j; j<1000000; j++){
	  pthread_mutex_lock( &mutex1);
	  i--;
	  pthread_mutex_unlock( &mutex1);
     }
}

int main(){
    pthread_t add_t;
    pthread_t sub_t;

    pthread_create(&add_t, NULL, add, NULL);
    pthread_create(&sub_t, NULL, sub, NULL);

    pthread_join(add_t, NULL);
    pthread_join(sub_t, NULL);

    printf("%d\n",i);
    return 0;
    
}
