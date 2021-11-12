package main

func buildMaxHeap(arr []int, n int) {
	for i := 1; i < n; i++ {
		//if child is bigger than parent
		if arr[i] > arr[(i-1)/2] {
			j := i
			// swap child and parent until
			// parent is smaller
			for arr[j] > arr[(j-1)/2] {
				swap(&arr[j], &arr[(j-1)/2])
				j = (j - 1) / 2
			}
		}
	}
}

func heapSort(arr []int, n int) {
	buildMaxHeap(arr, n)
	for i := n - 1; i > 0; i-- {
		// swap value of first indexed
		// with last indexed
		swap(&arr[0], &arr[i])
		// maintaining heap property
		// after each swapping
		j := 0
		index := 0

		for {
			index = 2*j + 1
			if index > i {
				break
			}
			// if left child is smaller than
			// right child point index variable
			// to right child
			if arr[index] < arr[index+1] && index < (i-1) {
				index++
			}
			// if parent is smaller than child
			// then swapping parent with child
			// having higher value
			if arr[j] < arr[index] && index < i {
				swap(&arr[j], &arr[index])
			}
			j = index
		}
	}
}

/*
// Driver Code to test above
int main()
{
int arr[] = {10, 20, 15, 17, 9, 21};
int n = sizeof(arr) / sizeof(arr[0]);

printf("Given array: ");
for (int i = 0; i < n; i++)
printf("%d ", arr[i]);

printf("\n\n");

heapSort(arr, n);

// print array after sorting
printf("Sorted array: ");
for (int i = 0; i < n; i++)
printf("%d ", arr[i]);

return 0;
}*/
