using System;
using System.Collections.Generic;
using System.ComponentModel;
using System.Data;
using System.Drawing;
using System.IO;
using System.Linq;
using System.Text;
using System.Windows.Forms;

namespace ProtocolTest
{
    public partial class Form1 : Form
    {
        public Form1()
        {
            InitializeComponent();
        }

        private void button1_Click(object sender, EventArgs e)
        {
            byte[] data = new byte[32+128];

            //产生一些随机数据当消息
            Random rand = new Random();
            rand.NextBytes(data);


            Protocol pt = new Protocol();

            byte[] sndData = pt.encap(data,data.Length);


            //写入到文件中去

            string filePath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "test.bin");
            FileStream fs = new FileStream(filePath, FileMode.Append);


            fs.Write(sndData, 0, sndData.Length);

            fs.Close();

            MessageBox.Show("发送OK");



        }

        private void button2_Click(object sender, EventArgs e)
        {
            string filePath = Path.Combine(AppDomain.CurrentDomain.BaseDirectory, "test.bin");
            FileStream fs = new FileStream(filePath, FileMode.Open);

            Protocol pt = new Protocol();

            byte[] data = new byte[16];

            while (true)
            {
               int length= fs.Read(data, 0, data.Length);

               if (length > 0)
               {
                   pt.parse(data, length);
               }
               if (length < data.Length)
               {
                   break;
               }
            }

            fs.Close();

        }
    }
}
