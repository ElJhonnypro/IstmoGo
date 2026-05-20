import { Button, StyleSheet } from "react-native";
import colors from "./colors";
import radius from "./radius";

export default StyleSheet.create({
    Input: {
        width: 300,
        padding: 15,
        backgroundColor: colors.Third,
        borderRadius: radius.Small
    },
    Label: {
        fontSize: 16,
        fontWeight: "bold",
        marginBottom: 5,
    },
    Button: {
        width: 300,
        alignContent: "center",
        alignItems: "center",
        padding: 15,
        backgroundColor: colors.secondary,
        borderRadius: radius.Small,
        fontSize: 16,
        fontWeight: "bold",
        color: colors.Third,
    },
    ButtonText: {
        fontSize: 16,
        fontWeight: "bold",
        color: colors.Third,
        textAlign: "center",
    },

    link: {
        color: colors.secondary,
        textDecorationLine: "underline",
    }
})