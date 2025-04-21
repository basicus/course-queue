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

public class KafkaStreamsJoinApp {
    public static void main(String[] args) throws Exception {

        Properties props = new Properties();
        props.put(StreamsConfig.APPLICATION_ID_CONFIG, "kafka-streams-aggregate");
        props.put(StreamsConfig.BOOTSTRAP_SERVERS_CONFIG, "localhost:29092,localhost:29093,localhost:29094,localhost:29095");
        props.put(StreamsConfig.DEFAULT_KEY_SERDE_CLASS_CONFIG, Serdes.String().getClass());
        props.put(StreamsConfig.DEFAULT_VALUE_SERDE_CLASS_CONFIG, Serdes.String().getClass());

        // Создайте StreamsBuilder экземпляр
        StreamsBuilder builder = new StreamsBuilder();

        // Таблица

        GlobalKTable<String, String> globalTable = builder.globalTable(
                "customers",
                Materialized.with(Serdes.String(), Serdes.String())
        );

        ValueJoiner<String, String, String> valueJoiner = (leftValue, rightValue) -> {
            PaymentEvent paymentEvent = new Gson().fromJson(String.valueOf(leftValue), PaymentEvent.class);
            paymentEvent.setCustomerName(rightValue);
            return new Gson().toJson(paymentEvent).toString();
        };


        KStream<String, String> sourceStream = builder.stream("payments");

        sourceStream
                .peek((key, value) -> System.out.println("In  >> key: " + key + ":\t" + value))
                .join(
                        globalTable,
                        (key, value) -> key,
                        valueJoiner

                )
                .peek((key, value) -> System.out.println("Out << key: " + key + ":\t" + value))
                .to("customer-payments", Produced.with(Serdes.String(), Serdes.String()));



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
