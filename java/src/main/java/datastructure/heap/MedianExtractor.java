package datastructure.heap;

import java.util.Objects;
import java.util.PriorityQueue;
import java.util.Queue;

public class MedianExtractor {

    private Queue<Integer> queue;

    private int counter;

    public MedianExtractor(int[] input) {
        Objects.requireNonNull(input);
        queue = new PriorityQueue<>();
        for (int i = 0; i < input.length/2+1; i++) {
            queue.offer(input[i]);
        }
        for (int i = input.length/2+1; i < input.length; i++) {
            Integer min = queue.peek();
            if (min < input[i]) {
                queue.poll();
                queue.offer(input[i]);
            }
        }
        counter = input.length;
    }

    public double getMedian() {
        if (counter % 2 == 0) {
            Integer min1 = queue.poll();
            Integer min2 = queue.poll();
            return (min1 + min2)/2f;
        }
        return queue.peek();
    }

    public static void main(String[] args) {
        int[] data = new int[]{3, 13, 7, 5, 21, 23, 39, 23, 40, 23, 14, 12, 23, 29, 56};
        MedianExtractor medianExtractor = new MedianExtractor(data);
        System.out.println(medianExtractor.getMedian());
    }
}
