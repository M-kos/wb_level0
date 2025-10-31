import { Button, Card, Flex, Input } from 'antd';
import { useState } from 'react';
import * as React from 'react';
import { useNavigate } from 'react-router-dom';

export const Home = () => {
  const navigate = useNavigate();
  const [uuid, setUuid] = useState('');

  const onChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setUuid(e.target.value);
  };

  return (
    <Card styles={{ body: { width: 400, height: 300 } }}>
      <Flex style={{ height: '100%' }} vertical gap={24} justify={'center'} align={'stretch'}>
        <Input
          placeholder={'Enter the order UUID'}
          size={'large'}
          value={uuid}
          onChange={onChange}
        />
        <Button
          htmlType={'submit'}
          type={'primary'}
          onClick={() => {
            navigate(`/order/${uuid}`);
          }}
        >
          {'Submit'}
        </Button>
      </Flex>
    </Card>
  );
};
