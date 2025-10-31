import { useLoaderData } from 'react-router-dom';
import type { Order } from '../../shared/api/order.types';
import { Card, Flex, Typography } from 'antd';

const { Text } = Typography;

export const OrderPage = () => {
  const data = useLoaderData<{ order: Order } | undefined>();

  if (!data?.order) {
    return <div>The order was not found</div>;
  }

  return (
    <Flex vertical style={{ width: '50%' }} gap={24}>
      <Flex gap={24}>
        <Card style={{ flex: 1 }} title={'Order'}>
          <Flex vertical>
            <Text>{`Order Uid: ${data.order.order_uid}`}</Text>
            <Text>{`Track Number: ${data.order.track_number}`}</Text>
            <Text>{`Entry: ${data.order.entry}`}</Text>
            <Text>{`Locale: ${data.order.locale}`}</Text>
            <Text>{`Internal Signature: ${data.order.internal_signature}`}</Text>
            <Text>{`Customer ID: ${data.order.customer_id}`}</Text>
            <Text>{`Delivery Service: ${data.order.delivery_service}`}</Text>
            <Text>{`Shardkey: ${data.order.shardkey}`}</Text>
            <Text>{`SM ID: ${data.order.order_uid}`}</Text>
            <Text>{`Date Created: ${data.order.order_uid}`}</Text>
            <Text>{`Oof Shard: ${data.order.order_uid}`}</Text>
          </Flex>
        </Card>
        <Card style={{ flex: 1 }} title={'Delivery'}>
          <Flex vertical>
            <Text>{`Name: ${data.order.delivery.name}`}</Text>
            <Text>{`Phone: ${data.order.delivery.phone}`}</Text>
            <Text>{`Zip: ${data.order.delivery.zip}`}</Text>
            <Text>{`City: ${data.order.delivery.city}`}</Text>
            <Text>{`Address: ${data.order.delivery.address}`}</Text>
            <Text>{`Region: ${data.order.delivery.region}`}</Text>
            <Text>{`Email: ${data.order.delivery.email}`}</Text>
          </Flex>
        </Card>
        <Card style={{ flex: 1 }} title={'Payment'}>
          <Flex vertical>
            <Text>{`Transaction: ${data.order.payment.transaction}`}</Text>
            <Text>{`Request ID: ${data.order.payment.request_id}`}</Text>
            <Text>{`Currency: ${data.order.payment.currency}`}</Text>
            <Text>{`Provider: ${data.order.payment.provider}`}</Text>
            <Text>{`Amount: ${data.order.payment.amount}`}</Text>
            <Text>{`Payment Dt: ${data.order.payment.payment_dt}`}</Text>
            <Text>{`Bank: ${data.order.payment.bank}`}</Text>
            <Text>{`Delivery Cost: ${data.order.payment.delivery_cost}`}</Text>
            <Text>{`Goods Total: ${data.order.payment.goods_total}`}</Text>
            <Text>{`Custom Fee: ${data.order.payment.transaction}`}</Text>
          </Flex>
        </Card>
      </Flex>
      <Flex gap={24} wrap>
        {data.order.items.map((item) => (
          <Card style={{ minWidth: 300 }} key={item.rid} title={'Item'}>
            <Flex vertical>
              <Text>{`Chrt ID: ${item.chrt_id}`}</Text>
              <Text>{`Track Number: ${item.track_number}`}</Text>
              <Text>{`Price: ${item.price}`}</Text>
              <Text>{`RID: ${item.rid}`}</Text>
              <Text>{`Name: ${item.name}`}</Text>
              <Text>{`Sale: ${item.sale}`}</Text>
              <Text>{`Size: ${item.size}`}</Text>
              <Text>{`Total Price: ${item.total_price}`}</Text>
              <Text>{`NM ID: ${item.nm_id}`}</Text>
              <Text>{`Brand: ${item.brand}`}</Text>
              <Text>{`Status: ${item.status}`}</Text>
            </Flex>
          </Card>
        ))}
      </Flex>
    </Flex>
  );
};
