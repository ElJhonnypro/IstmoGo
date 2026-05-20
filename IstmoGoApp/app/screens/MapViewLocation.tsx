import { useEffect, useState } from "react";
import { Dimensions, Pressable, ScrollView, Text, View } from "react-native";
import * as Location from "expo-location";
import MapView, { Marker } from "react-native-maps";
import Inputs from "../Styles/Components/Inputs";
import { NavigationProp, useNavigation } from "@react-navigation/native";
import { RootStackParamList } from "../types/navigation";

export default function MapViewLocation({ route }: any) {
  const { width, height } = Dimensions.get("window");
  const navigation = useNavigation<NavigationProp<RootStackParamList>>();
  const ASPECT_RATIO = width / height;
  const LATITUDE_DELTA = 0.0007;
  const LONGITUDE_DELTA = LATITUDE_DELTA * ASPECT_RATIO;

  const { goBack } = route.params;
  const [markerLocation, setMarkerLocation] = useState(null);
  const [location, setLocation] = useState({
    latitude: 0,
    longitude: 0,
  });
  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();

      if (status !== "granted") {
        alert("Permiso de ubicación denegado");
        return;
      }

      const loc = await Location.getCurrentPositionAsync({});

      setLocation({
        latitude: loc.coords.latitude,
        longitude: loc.coords.longitude,
      });
    })();
  }, []);

  const handleMapPress = (event: any) => {
    // Get the coordinates from the click event
    const { coordinate } = event.nativeEvent;
    setMarkerLocation(coordinate);
  };

  if (location.latitude === 0) {
    return (
      <View>
        <Text>Hello help me please</Text>
      </View>
    );
  }

  return (
    <View
      style={{
        display: "flex",
        flexDirection: "column",
        flex: 1,
        height: "100%",
      }}
    >
      <MapView
        style={{ flex: 1 }}
        initialCamera={{
          center: {
            latitude: location.latitude,
            longitude: location.longitude,
          },
          heading: 0, // 0 = Norte, 90 = Este, 180 = Sur, 270 = Oeste
          pitch: 25, // 0 = Vista plana desde arriba
          zoom: 35, // Zoom de calle
          altitude: 1000,
        }}
        onPress={handleMapPress}
      >
        {markerLocation && (
          <Marker coordinate={markerLocation} title="Picked Location" />
        )}
      </MapView>
      {markerLocation && (
        <Pressable
          onPress={() => {
            navigation.navigate("Register", markerLocation);
          }}
          style={[
            Inputs.Button,
            {
              position: "absolute",
              bottom: 100,
              left: 50,
              right: 50,
              alignItems: "center",
              zIndex: 1,

              shadowColor: "#1d2823ff",
              shadowOffset: { width: 0, height: 4 },
              shadowOpacity: 0.25,
              shadowRadius: 10,
              elevation: 15,
            },
          ]}
        >
          <Text style={Inputs.ButtonText}>Aceptar</Text>
        </Pressable>
      )}
    </View>
  );
}
