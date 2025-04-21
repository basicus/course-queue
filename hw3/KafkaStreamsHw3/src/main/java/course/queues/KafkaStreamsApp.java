package course.queues;

import org.apache.kafka.streams.KafkaStreams;
import org.apache.kafka.streams.StreamsBuilder;
import org.apache.kafka.streams.StreamsConfig;
import org.apache.kafka.common.serialization.Serdes;
import org.apache.kafka.streams.kstream.*;
import com.google.gson.Gson;


import java.time.Duration;
import java.util.Properties;
import java.util.concurrent.CountDownLatch;

public class KafkaStreamsApp {
    public static void main(String[] args) throws Exception {

        Properties props = new Properties();
        props.put(StreamsConfig.APPLICATION_ID_CONFIG, "kafka-streams-aggregate");
        props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:29092,localhost:29093,localhost:29094,localhost:29095");
        props.put(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG, Serdes.String().getClass());
        props.put(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG, Serdes.String().getClass());

        // Создайте StreamsBuilder экземпляр
        StreamsBuilder builder = new StreamsBuilder();

        // Определите топологию Kafka Streams приложения


        KStream<String, String> sourceStream = builder.stream("payments");

        sourceStream.peek((key, value) -> System.out.println("In  >> key: " + key + ":\t" + value))
                .mapValues(value -> {
                    PaymentEvent paymentEvent = new Gson().fromJson(String.valueOf(value), PaymentEvent.class);
                    return paymentEvent.getAmount();
                })
                .groupByKey()
                .windowedBy(
                        TimeWindows.ofSizeAndGrace(Duration.ofSeconds(15),
                                Duration.ofSeconds(1))
                )
                .aggregate(
                        () -> 0.0,
                        (key, value, total) -> (total + value)/2,
                        Materialized.with(Serdes.String(), Serdes.Double())
                )
                .toStream()
                .peek((key, value) -> System.out.println("Out << key: " + key + ":\t" + value + " (" + (key != null ? key.getClass().getName() : "-") + ", " + (value != null ? value.getClass().getName() : "-") + ")"))
                .to("payment-amounts-sum-avg");

        // Создайте KafkaStreams экземпляр
        try (KafkaStreams kafkaStreams = new KafkaStreams(builder.build(), props)) {
            final CountDownLatch shutdownLatch = new CountDownLatch(1);
            Runtime.getRuntime().addShutdownHook(new Thread(() -> {
                kafkaStreams.close(Duration.ofSeconds(10));
                shutdownLatch.countDown();
            }));
            try {
                kafkaStreams.start();
                shutdownLatch.await();
            } catch (Throwable e) {
                System.exit(1);
            }
        }
        System.exit(0);
    }
}
