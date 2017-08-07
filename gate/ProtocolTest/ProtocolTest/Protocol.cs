using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace ProtocolTest
{
    class Protocol
    {

        int frameLength = 0;

        byte[] frameBuffer=new byte[64*1024];

        bool met10 = false;


        /// <summary>
        /// 解析数据
        /// </summary>
        /// <param name="data"></param>
        /// <param name="length"></param>
        public void parse(byte[] data, int length)
        {
            for (int i = 0; i < length; i++)
            {
                switch (data[i])
                {
                    case 0x10:
                        {
                            if (frameLength == 0)
                            {
                                frameBuffer[frameLength++] = 0x10;

                            }
                            else if (frameLength > 1)
                            { 
                                if(!met10)
                                {
                                    frameBuffer[frameLength++] = 0x10;
                                }
                            }

                            met10 = !met10;
                        }
                        break;
                    case 0x02:
                        {
                            if (met10) //met a frame header
                            {
                                frameBuffer[0] = 0x10;
                                frameBuffer[1] = 0x02;
                                frameLength = 2;
                               
                            }
                            else if (frameLength > 1)
                            {
                                frameBuffer[frameLength++] = 0x02;
                            }

                            met10 = false;

                        }
                        break;
                    case 0x03:
                        {
                            if (met10) //met a frame tailer
                            {
                                frameBuffer[frameLength++] = 0x03;

                                //做长度和校验和验证

                                if (frameLength >=8) //2(帧头)+2(长度)+N(数据)+2(校验和)+2(帧尾)
                                {
                                    UInt16 fLen = BitConverter.ToUInt16(frameBuffer, 2);

                                    ushort sum = 0;
                                    if (fLen == (frameLength - 2 - 2 - 2))//做长度的校验
                                    {
                                        sum = 0;
                                        for (int j = 0; j < fLen - 2; j++)
                                        {
                                            sum += frameBuffer[4 + j];
                                        }

                                        byte checkLow = frameBuffer[2 + fLen];
                                        byte checkHigh = frameBuffer[2 + fLen + 1];

                                        if (((sum & 0xff) == checkLow) && ((sum >> 8) & 0xff) == checkHigh) //校验和通过
                                        {
                                            //这个时候，数据帧完全OK

                                            Console.WriteLine(" message OK");
                                        }

                                    }
                                }
                                this.frameLength = 0; //开始新的一帧的获取
                               
                            }
                            else if (frameLength > 1)
                            {
                                frameBuffer[frameLength++] = 0x03;
                            }
                            met10 = false;
                        }
                        break;
                    default:
                        {
                            if (frameLength > 1)
                            {
                                frameBuffer[frameLength++] = data[i];
                            }
                        }
                        break;
                }
            }
        }



        /// <summary>
        /// 封装发送的数据
        /// </summary>
        /// <param name="data">这个就是要发送的消息(消息头+消息体)</param>
        /// <param name="length"></param>
        /// <returns></returns>
        public byte[] encap(byte[] data, int length)
        {
            List<byte> listSnd=new List<byte>();


            //计算和校验

            UInt16 sum = 0;

            for (int i = 0; i < length; i++)
            {
                sum += data[i];
            }

            UInt16 fLen = (UInt16)(length + 2);

            //帧头

            listSnd.Add(0x10);
            listSnd.Add(0x02);

            //长度
            byte lenLow = (byte)(fLen & 0xff);
            listSnd.Add(lenLow);
            if (lenLow == 0x10)
            {
                listSnd.Add(0x10);
            }
            byte lenHigh = (byte)((fLen>>8) & 0xff);
            listSnd.Add(lenHigh);
            if (lenHigh == 0x10)
            {
                listSnd.Add(0x10);
            }

            //数据

            for (int i = 0; i < length; i++)
            {
                listSnd.Add(data[i]);
                if (data[i] == 0x10)
                {
                    listSnd.Add(0x10);
                }
            }

            //校验和

            byte sumLow = (byte)(sum & 0xff);
            listSnd.Add(sumLow);
            if (sumLow == 0x10)
            {
                listSnd.Add(0x10);
            }
            byte sumHigh = (byte)((sum >> 8) & 0xff);
            listSnd.Add(sumHigh);
            if (sumHigh == 0x10)
            {
                listSnd.Add(0x10);
            }

            //帧尾

            listSnd.Add(0x10);
            listSnd.Add(0x03);

            return listSnd.ToArray();
        }


    }
}
