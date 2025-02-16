import { useEffect, useState } from 'react';
import { View, Text, ActivityIndicator, Button } from 'react-native';
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

  if (!token) {
    return <ActivityIndicator />;
  }

  return (
    <View>
      <Text>ログイン成功！</Text>
      <Text>アクセストークン: {token}</Text>
      <Button title = "ログアウト" onPress={handle}></Button>
    </View>
  );
}
