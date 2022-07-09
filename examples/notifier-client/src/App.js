import logo from './logo.svg';
import './App.css';
import React, {useState, useEffect} from 'react';
import { v4 as uuidv4 } from 'uuid';

import { NotifierClient } from '@circadence/web/atlas/notifier/v1/notifier_grpc_web_pb'
import { ConnectionRequest, Notification } from '@circadence/web/atlas/notifier/v1/notifier_pb'

const client = new NotifierClient("");
const url_string = window.location.href;
const url = new URL(url_string); 
const user_id = url.searchParams.get("id");
const client_id = uuidv4();

function App() {
  const [notifications, setNotification] = useState([]);

  const getNotifications = () => {

    console.log("user_id: ", user_id);
    console.log("client_id: ", client_id);

    const  req = new ConnectionRequest()
      .setClientId(client_id)
      .setUserId(user_id);

    const stream = client.connect(req)

    stream.on('data', function(res)  {
      let data = res.getData()
      console.log("notification type: ", res.getNotificationType())
      console.log("notifiction data: ", JSON.parse(data))
    })

    stream.on('status', function(res) {
      console.log("status res: ", res)
    })

    stream.on('end', function(res) {
      console.log("end res: ", "ended")
    })
  }

  useEffect(() => {
    getNotifications();
  })

  return (
    <div className="App">
      
    </div>
  );
}

export default App;
