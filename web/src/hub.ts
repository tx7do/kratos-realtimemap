export interface PositionsDto {
  positions: PositionDto[];
}

export interface PositionDto {
  vehicleId: string;
  longitude: number;
  latitude: number;
  heading: number;
  speed: number;
  doorsOpen: boolean;
}

export interface WebsocketProto {
  eventId: string;
  payload: string;
}

export interface GeoPoint {
  longitude: number;
  latitude: number;
}

export interface Viewport {
  southWest: GeoPoint;
  northEast: GeoPoint;
}

export interface UpdateViewport {
  viewport: Viewport;
}

export interface Notification {
  message: string;
}

export interface HubConnection {
  setViewport(swLng: number, swLat: number, neLng: number, neLat: number);

  onPositions(callback: (positions: PositionDto[]) => void);

  onNotification(callback: (notification: string) => void);

  disconnect(): Promise<void>;
}

function ByteBufferToObject(buff) {
  const enc = new TextDecoder('utf-8');
  const uint8Array = new Uint8Array(buff);
  const decodedString = enc.decode(uint8Array);
  // console.log(decodedString);
  return JSON.parse(decodedString);
}

function StringToArrayBuffer(str) {
  const buf = new ArrayBuffer(str.length * 2); // 每个字符占用2个字节
  const bufView = new Uint16Array(buf);
  for (let i = 0, strLen = str.length; i < strLen; i++) {
    bufView[i] = str.charCodeAt(i);
  }
  return buf;
}

class WebsocketConnect implements HubConnection {
  private connection: WebSocket;
  private onPositionsCallback?: (positions: PositionDto[]) => void;
  private onNotificationCallback?: (notification: string) => void;

  constructor() {

    const wsURL = `ws://localhost:7700/`;
    this.connection = new WebSocket(wsURL);
    this.connection.binaryType = 'arraybuffer';
    this.connection.onopen = this.onWebsocketOpen;
    this.connection.onerror = this.onWebsocketError;
    this.connection.onmessage = this.onWebsocketMessage;
    this.connection.onclose = this.onWebsocketClose;
  }

  onWebsocketOpen(e) {
    console.log('ws连接成功', e);
  }

  onWebsocketError(e) {
    console.error('ws错误', e);
  }

  onWebsocketMessage(evt) {
    const proto = ByteBufferToObject(evt.data);
    // console.log(proto);
    const data = JSON.parse(proto['payload']);
    // console.log(data);

    const eventId = proto['event_id'];
    if (eventId == 'positions') {
      if (this.onPositionsCallback != null) {
        this.onPositionsCallback(data);
      }
    } else if (eventId == 'notification') {
      if (this.onNotificationCallback != null) {
        this.onNotificationCallback(data);
      }
    }
  }

  onWebsocketClose(e) {
    console.log('ws连接关闭', e);
  }

  sendMessage(eventId, data) {
    const x: WebsocketProto = {
      eventId: eventId,
      payload: data,
    }
    this.connection.send(StringToArrayBuffer(x));
  }

  setViewport(swLng: number, swLat: number, neLng: number, neLat: number) {
    const x: Viewport = {
      southWest: {
        longitude: swLng,
        latitude: swLat,
      },
      northEast: {
        longitude: neLng,
        latitude: neLat,
      },
    };
    this.sendMessage('viewport', x);
  }

  onPositions(callback) {
    this.onPositionsCallback = callback;
  }

  onNotification(callback) {
    this.onNotificationCallback = callback;
  }

  async disconnect() {
    await this.connection.close(0);
  }
}

export const connectToHub = WebsocketConnect;
