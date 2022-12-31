import './App.css';
import React, { useState, useEffect } from 'react';
import { v4 as uuidv4 } from 'uuid';

import { NotifierClient } from '@api/web/core/notifier/v1/notifier_grpc_web_pb';
import {
  ConnectionRequest,
  Notification,
} from '@api/web/core/notifier/v1/notifier_pb';

const client = new NotifierClient('');
const url_string = window.location.href;
const url = new URL(url_string);
const actor_id = url.searchParams.get('id');
const client_id = uuidv4();
const transaction_id = uuidv4();

function App() {
  const [notifications, setNotification] = useState([]);

  const getNotifications = () => {
    console.log('actor_id: ', actor_id);
    console.log('client_id: ', client_id);
    console.log('transaction_id: ', transaction_id)

    const req = new ConnectionRequest()
      .setClientId(client_id)
      .setActorId(actor_id);

    var meta = {'transaction_id': transaction_id};
    const stream = client.connect(req, meta);

    stream.on('data', function (res) {
      let data = res.getData();
      console.log('notification type: ', res.getNotificationType());
      console.log('notifiction data: ', JSON.parse(data));
    });

    stream.on('status', function (res) {
      console.log('status res: ', res);
    });

    stream.on('end', function (res) {
      console.log('end res: ', 'ended');
    });
  };

  useEffect(() => {
    getNotifications();
  });

  return <div className="App"></div>;
}

export default App;
