import time
import pandas as pd
import streamlit as st
import plotly.express as px
from utils import control_fan, get_sensor_data

st.set_page_config(
    page_title="ESP32 Dashboard",
    page_icon="🌡️",
)

st.title("Dashboard")

st.subheader("Project Overview")
st.write(
    """
    This project is an IoT-based monitoring system that tracks temperature and humidity\
        levels in real-time using a DHT11 sensor connected to an ESP32 microcontroller.
    The data is collected by an API built with Go, stored in a PostgreSQL database, and\
        presented through this interactive dashboard built with Streamlit.

    Users can monitor live temperature and humidity trends, view historical data, and\
        control an external fan directly from the dashboard, making this system ideal
    for environments where maintaining specific temperature and humidity conditions\
        is crucial.
    """
)

if "fan_state" not in st.session_state:
    st.session_state.fan_state = "off"


def fetch_data() -> pd.DataFrame:
    sensor_data = get_sensor_data()
    if sensor_data is not None and not sensor_data.empty:
        sensor_data["Timestamp"] = pd.to_datetime(
            sensor_data["Timestamp"]
        ) - pd.Timedelta(hours=3)
        sensor_data["Date"] = sensor_data["Timestamp"].dt.date
        sensor_data["Time"] = sensor_data["Timestamp"].dt.strftime("%H:%M")
        return sensor_data
    else:
        return pd.DataFrame(
            columns=["ID", "Temperature", "Humidity", "Timestamp", "Date", "Time"]
        )


with st.sidebar:
    st.subheader("Fan Control")
    if st.button("Toggle Fan"):
        new_state = "off" if st.session_state.fan_state == "on" else "on"
        control_fan(new_state)
        st.session_state.fan_state = new_state

    if st.session_state.fan_state == "on":
        st.success("Fan is ON")
    else:
        st.error("Fan is OFF")

    st.subheader("Date Filter")
    data = fetch_data()
    available_dates = data["Date"].unique()
    selected_date = st.selectbox("Select Date", options=available_dates)

    st.subheader("Time Filter")
    hours = [f"{hour:02d}:00" for hour in range(24)]
    start_time = st.selectbox("Start Time", options=hours, index=0)
    end_time = st.selectbox("End Time", options=hours, index=23)

if selected_date is not None:
    data = data[
        (data["Date"] == selected_date)
        & (data["Time"] >= start_time)
        & (data["Time"] <= end_time)
    ]

st.subheader("Temperature Plot")
fig_temp = px.line(data, x="Timestamp", y="Temperature", title="Temperature over Time")
fig_temp.update_yaxes(rangemode="tozero", range=[data["Temperature"].max() - 5, data["Temperature"].max() + 5])
st.plotly_chart(fig_temp)

humidity_col, df_col = st.columns([2, 2])

with humidity_col:
    st.subheader("Humidity Plot")
    fig_humidity = px.line(data, x="Timestamp", y="Humidity", title="Humidity over Time")
    fig_humidity.update_yaxes(rangemode="tozero", range=[data["Humidity"].min() - 5, data["Humidity"].max() + 5])
    st.plotly_chart(fig_humidity)

with df_col:
    st.subheader("Sensor Data")
    st.dataframe(data[["ID", "Temperature", "Humidity", "Date", "Time"]])

while True:
    data = fetch_data()

    if selected_date is not None:
        data = data[
            (data["Date"] == selected_date)
            & (data["Time"] >= start_time)
            & (data["Time"] <= end_time)
        ]

    fig_temp = px.line(data, x="Timestamp", y="Temperature", title="Temperature over Time")
    fig_temp.update_yaxes(rangemode="tozero", range=[18, data["Temperature"].max() + 5])
    fig_humidity = px.line(data, x="Timestamp", y="Humidity", title="Humidity over Time")
    fig_humidity.update_yaxes(rangemode="tozero", range=[18, data["Humidity"].max() + 5])

    st.plotly_chart(fig_temp)
    st.plotly_chart(fig_humidity)
    st.dataframe(data[["ID", "Temperature", "Humidity", "Date", "Time"]])

    time.sleep(5)
    st.rerun()
