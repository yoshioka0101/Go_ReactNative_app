import { useEffect, useState } from 'react';
import { View, Text, ActivityIndicator, Button, StyleSheet } from 'react-native';
import { useLocalSearchParams } from 'expo-router';

export default function Dashboard() {
  const params = useLocalSearchParams();
  const [token, setToken] = useState<string | null>(null);

  useEffect(() => {
    if (params.token) {
      setToken(params.token as string);
      localStorage.setItem('accessToken', params.token as string);
    }
  }, [params]);

  const handleLogout = () => {
    localStorage.removeItem('accessToken');
    setToken(null);
    alert("ログアウトしました！！！！");
    window.location.href = "/";
  }

  if (!token) {
    return <ActivityIndicator />;
  }

  return (
    <View>
      <Text>ログイン成功！</Text>
      <View style={styles.buttonContainer}>
        <Button title = "ログアウト" onPress={handleLogout}></Button>
      </View>
    </View>
  );
}
const styles = StyleSheet.create({
    buttonContainer: {
     marginTop: 16,
     alignItems: 'center',
    },
});