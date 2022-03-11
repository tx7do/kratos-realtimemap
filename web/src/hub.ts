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

export interface HubConnection {
  setViewport(swLng: number, swLat: number, neLng: number, neLat: number);

  onPositions(callback: (positions: PositionsDto) => void);

  onNotification(callback: (notification: string) => void);

  disconnect(): Promise<void>;
}

class WebsocketConnect implements HubConnection {
  private connection: WebSocket;

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
    const enc = new TextDecoder('utf-8');
    const uint8Array = new Uint8Array(evt.data);
    console.log(enc.decode(uint8Array));
    const decodedString = enc.decode(uint8Array);
    const data = JSON.parse(decodedString);
    console.log(data);
  }

  onWebsocketClose(e) {
    console.log('ws连接关闭', e);
  }

  sendMessage(data) {
    this.connection.send(data);
  }


  setViewport(swLng: number, swLat: number, neLng: number, neLat: number) {
    this.sendMessage({ swLng, swLat, neLng, neLat });
  }

  onPositions(callback) {
    callback();
  }

  onNotification(callback) {
    callback();
  }

  async disconnect() {
    await this.connection.close(0);
  }
}

export const connectToHub = WebsocketConnect;
