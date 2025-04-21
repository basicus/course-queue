package course.queues;

import com.google.gson.annotations.SerializedName;
import java.time.LocalDateTime;

public class PaymentEvent {
    @SerializedName("payment_id")
    private Long paymentId;

    @SerializedName("timestamp")
    private String timestamp;

    @SerializedName("amount")
    private Double amount;

    @SerializedName("status")
    private String status;

    @SerializedName("customer_id")
    private Long customerId;

    @SerializedName("customer_name")
    private String customerName;

    // Getters and setters
    public Long getPaymentId() {
        return paymentId;
    }

    public void setPaymentId(Long paymentId) {
        this.paymentId = paymentId;
    }

    public String getTimestamp() {
        return timestamp;
    }

    public void setTimestamp(String timestamp) {
        this.timestamp = timestamp;
    }

    public Double getAmount() {
        return amount;
    }

    public void setAmount(Double amount) {
        this.amount = amount;
    }

    public String getStatus() {
        return status;
    }

    public void setStatus(String status) {
        this.status = status;
    }

    public Long getCustomerId() {
        return customerId;
    }

    public void setCustomerId(Long customerId) {
        this.customerId = customerId;
    }

    public String getCustomerName() {
        return customerName;
    }

    public void setCustomerName(String customerName) {
        this.customerName = customerName;
    }
}
