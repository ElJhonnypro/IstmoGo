import { View, Text, Image, TextInput, Pressable } from "react-native";
import { useNavigation } from "@react-navigation/native";
import { styles as LStyles } from "../Styles/Components/LoginStyles";
import { styles as BStyles } from "../Styles/General/Backgrounds";
import radius from "../Styles/Components/radius";
import Inputs from "../Styles/Components/Inputs";
import colors from "../Styles/Components/colors";

export default function Login() {
  const navigation = useNavigation();

  function handleLogin() {
    // Aquí puedes agregar la lógica para manejar el inicio de sesión
  }

  return (
    <View style={[BStyles.main, LStyles.container]}>
      <View
        style={[
          BStyles.black,
          {
            borderRadius: radius.Normal,
            padding: 5,
            margin: 5,
            height: 574,
            width: 346,
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "flex-start",
          },
        ]}
      >
        <Image
          style={[
            LStyles.logo,
            {
              flex: 0,
            },
          ]}
          source={require("../../assets/images/LogoHorizontal.png")}
        />

        <View style={LStyles.Form}>
          <TextInput
            style={Inputs.Input}
            placeholderTextColor={colors.Fourth}
            placeholder="Correo electrónico / Usuario"
            keyboardType="email-address"
            autoCapitalize="none"
            autoCorrect={false}
            // Autocompletado:
            textContentType="emailAddress" // iOS
            autoComplete="email" // Android
          />
          <TextInput
            style={Inputs.Input}
            placeholderTextColor={colors.Fourth}
            placeholder="Contraseña"
            secureTextEntry
            textContentType="password" // iOS
            autoComplete="password" // Android
          />
          <Pressable
            style={({ pressed }) => [
              Inputs.Button,
              {
                backgroundColor: pressed
                  ? colors.secondary + "CC"
                  : colors.secondary,

                color: pressed ? colors.Third + "CC" : colors.Third,
              },
            ]}
            onPress={handleLogin}
          >
            <Text style={Inputs.ButtonText}>Iniciar sesión</Text>
          </Pressable>

          <Pressable onPress={() => navigation.navigate("Register" as never)}>
            <Text style={Inputs.link}>¿No tienes una cuenta? Regístrate</Text>
          </Pressable>
        </View>
      </View>
    </View>
  );
}
