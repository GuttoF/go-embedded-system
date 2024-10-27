#include <WiFi.h>
#include <HTTPClient.h>
#include <DHT.h>

#define DHTPIN 5
#define DHTTYPE DHT11
#define LED_PIN 18

// Replace with your network credentials
const char* ssid = "";
const char* password = "";
iconst char* serverTemp = "";
const char* serverFan = "";

DHT dht(DHTPIN, DHTTYPE);

void connectToWiFi() {
	WiFi.begin(ssid, password);
	Serial.print("Connecting to WiFi");
	while (WiFi.status() != WL_CONNECTED) {
		delay(500);
		Serial.print(".");
	}
	Serial.println("Connected to WiFi");
}

float readTemperature() {
	float temp = dht.readTemperature();
	if (isnan(temp)) {
	Serial.println("Failed to read temperature from DHT sensor!");
	}
	return temp;
}


float readHumidity() {
	float hum = dht.readHumidity();
	if (isnan(hum)) {
		Serial.println("Failed to read humidity from DHT sensor!");
	}
	return hum;
}


void sendTemperatureData(float temperature, float humidity) {
	if (WiFi.status() == WL_CONNECTED) {
		HTTPClient http;
		http.begin(serverTemp);


		String jsonPayload = "{\"temperature\":" + String(temperature) + ",\"humidity\":" + String(humidity) + "}";
		http.addHeader("Content-Type", "application/json");

		int httpResponseCode = http.POST(jsonPayload);

	if (httpResponseCode > 0) {
		Serial.print("Temperature and Humidity Data Sent. Response code: ");
		Serial.println(httpResponseCode);
	} else {
		Serial.print("Error on sending POST: ");
		Serial.println(httpResponseCode);
	}
		http.end();
	} else {
		Serial.println("Disconnected from WiFi");
	}
}


void controlFan() {
	if (WiFi.status() == WL_CONNECTED) {
	HTTPClient http;
	http.begin(serverFan);
	http.addHeader("Content-Type", "application/json");


	int httpResponseCode = http.POST("");

	if (httpResponseCode > 0) {
		String response = http.getString();
		Serial.println("Fan Control Response: " + response);


		if (response.indexOf("\"fan\":\"on\"") > 0) {
			digitalWrite(LED_PIN, HIGH);
			Serial.println("Fan ON");
		} else if (response.indexOf("\"fan\":\"off\"") > 0) {
			digitalWrite(LED_PIN, LOW);
			Serial.println("Fan OFF");
		}
		} else {
			Serial.print("Error on fan control POST: ");
			Serial.println(httpResponseCode);
		}
	http.end();
	}
}


void setup() {
	Serial.begin(115200);
	pinMode(LED_PIN, OUTPUT);
	digitalWrite(LED_PIN, LOW);

	dht.begin();
	connectToWiFi();
}

void loop() {
	float temperature = readTemperature();
	float humidity = readHumidity();

	if (!isnan(temperature) && !isnan(humidity)) {
		sendTemperatureData(temperature, humidity);
		controlFan();
	}

	delay(5000);
}