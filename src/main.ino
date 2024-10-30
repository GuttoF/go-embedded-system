#include <WiFi.h>
#include <HTTPClient.h>
#include <DHT.h>

// Wi-Fi configuration
const char* ssid = "";
const char* password = "";

#define DHTPIN 26
#define FAN_PIN 2
#define DHTTYPE DHT11

DHT dht(DHTPIN, DHTTYPE);
String apiUrl = ":3000";  // Replace

void connectWiFi() {
  WiFi.begin(ssid, password);
  Serial.print("Connecting to WiFi");
  while (WiFi.status() != WL_CONNECTED) {
    delay(1000);
    Serial.print(".");
  }
  Serial.println("\nConnected to WiFi!");
}

void initHardware() {
  dht.begin();
  pinMode(FAN_PIN, OUTPUT);
}

void checkFanStatus() {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String serverPath = apiUrl + "/fan/status";

    http.begin(serverPath.c_str());
    int httpResponseCode = http.GET();

    if (httpResponseCode == 200) {
      String response = http.getString();
      if (response.indexOf("\"state\":\"on\"") > -1) {
        digitalWrite(FAN_PIN, HIGH);
        Serial.println("Fan turned ON via dashboard");
      } else if (response.indexOf("\"state\":\"off\"") > -1) {
        digitalWrite(FAN_PIN, LOW);
        Serial.println("Fan turned OFF via dashboard");
      }
    } else {
      Serial.print("Error checking fan status: ");
      Serial.println(httpResponseCode);
    }
    http.end();
  } else {
    Serial.println("WiFi disconnected, attempting to reconnect...");
    connectWiFi();
  }
}

void sendTemperatureData(float temperature, float humidity) {
  if (WiFi.status() == WL_CONNECTED) {
    HTTPClient http;
    String serverPath = apiUrl + "/temperature";

    http.begin(serverPath.c_str());
    http.addHeader("Content-Type", "application/json");

    String jsonPayload = "{\"temperature\":" + String(temperature) + ",\"humidity\":" + String(humidity) + "}";
    int httpResponseCode = http.POST(jsonPayload);

    if (httpResponseCode > 0) {
      Serial.print("Response code: ");
      Serial.println(httpResponseCode);
      Serial.print("API response: ");
      Serial.println(http.getString());
    } else {
      Serial.print("Error sending data: ");
      Serial.println(httpResponseCode);
    }
    http.end();
  }
}

void readSensorData(float &temperature, float &humidity) {
  temperature = dht.readTemperature();
  humidity = dht.readHumidity();

  if (isnan(temperature) || isnan(humidity)) {
    Serial.println("Failed to read from DHT sensor");
  } else {
    Serial.print("Temperature: ");
    Serial.print(temperature);
    Serial.print(" °C, Humidity: ");
    Serial.print(humidity);
    Serial.println(" %");
  }
}

void setup() {
  Serial.begin(115200);
  connectWiFi();
  initHardware();
}

void loop() {
  float temperature = 0.0;
  float humidity = 0.0;

  readSensorData(temperature, humidity);
  sendTemperatureData(temperature, humidity);
  checkFanStatus();

  delay(5000);
}