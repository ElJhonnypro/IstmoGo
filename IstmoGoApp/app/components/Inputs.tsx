import { TextInput, TextInputProps } from "react-native";
import Inputs from "../Styles/Components/Inputs";
import colors from "../Styles/Components/colors";

type Props = TextInputProps & {
  onChangeText?: (text: string) => void;
};

export default function JInputText({ style, onChangeText, ...props }: Props) {
  return (
    <TextInput
      style={[Inputs.Input, style]}
      placeholderTextColor={colors.Fourth}
      onChangeText={onChangeText}
      {...props}
    />
  );
}
