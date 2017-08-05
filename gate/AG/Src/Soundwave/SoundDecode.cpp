#include "SoundDecode.h"
#include <alsa/asoundlib.h>
#include <fftw3.h>

SoundDecode::SoundDecode()
{
    this->isRunning=false;
}


void SoundDecode::decode()
{
#ifdef __arm__
       char * soundcard="hw:1,0";
#else
       char * soundcard="hw:2,0";
#endif

    int err;
    snd_pcm_t *capture_handle;
    snd_pcm_hw_params_t *hw_params;
    snd_pcm_format_t format = SND_PCM_FORMAT_S16_LE;

    char *buffer;
    int buffer_frames = 48*20;
    unsigned int sampleRate = 48000;

    int channelNum=1;



    uint threshold=10000;


    int freqS0=1100;
    int freqS1=freqS0+63*100;

    fftw_complex *din,*out;
    fftw_plan p;
    int rate=10;
    int N=sampleRate/rate;
    din  = (fftw_complex*) fftw_malloc(sizeof(fftw_complex) * N);
    out = (fftw_complex*) fftw_malloc(sizeof(fftw_complex) * N);

    p = fftw_plan_dft_1d(N, din, out, FFTW_FORWARD,FFTW_ESTIMATE);


    if ((err = snd_pcm_open (&capture_handle, soundcard, SND_PCM_STREAM_CAPTURE, 0)) < 0) {
        fprintf (stderr, "cannot open audio device %s (%s)\n",
                 soundcard,
                snd_strerror (err));
        exit (1);
    }


     std::cout<<"audio interface opened\n"<<std::endl;


    if ((err = snd_pcm_hw_params_malloc (&hw_params)) < 0) {
        fprintf (stderr, "cannot allocate hardware parameter structure (%s)\n",
                 snd_strerror (err));
        exit (1);
    }

    std::cout<<"hw_params allocated\n"<<std::endl;

      if ((err = snd_pcm_hw_params_any (capture_handle, hw_params)) < 0) {
        fprintf (stderr, "cannot initialize hardware parameter structure (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

     std::cout<<"hw_params initialized\n"<<std::endl;

      if ((err = snd_pcm_hw_params_set_access (capture_handle, hw_params, SND_PCM_ACCESS_RW_INTERLEAVED)) < 0) {
        fprintf (stderr, "cannot set access type (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

     std::cout<<"hw_params access setted\n"<<std::endl;

      if ((err = snd_pcm_hw_params_set_format (capture_handle, hw_params, format)) < 0) {
        fprintf (stderr, "cannot set sample format (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

     std::cout<<"hw_params format setted\n"<<std::endl;

      if ((err = snd_pcm_hw_params_set_rate_near (capture_handle, hw_params, &sampleRate, 0)) < 0) {
        fprintf (stderr, "cannot set sample rate (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

    std::cout<<"hw_params rate setted\n"<<std::endl;

      if ((err = snd_pcm_hw_params_set_channels (capture_handle, hw_params, channelNum)) < 0) {
        fprintf (stderr, "cannot set channel count (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

       std::cout<<"hw_params channels setted\n"<<std::endl;

      if ((err = snd_pcm_hw_params (capture_handle, hw_params)) < 0) {
        fprintf (stderr, "cannot set parameters (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

       std::cout<<"hw_params setted\n"<<std::endl;

      snd_pcm_hw_params_free (hw_params);

       std::cout<<"hw_params freed\n"<<std::endl;

      if ((err = snd_pcm_prepare (capture_handle)) < 0) {
        fprintf (stderr, "cannot prepare audio interface for use (%s)\n",
                 snd_strerror (err));
        exit (1);
      }

      fprintf(stdout, "audio interface prepared\n");


      int bufferSize=buffer_frames * snd_pcm_format_width(format) / 8 * channelNum;
      buffer = (char*)malloc(bufferSize);

      fprintf(stdout, "buffer allocated\n");


      while (this->isRunning) {

          if ((err = snd_pcm_readi (capture_handle, buffer, buffer_frames)) != buffer_frames) {

            break;
          }

          for(int i=0;i<buffer_frames;i++)
          {
                din[i][0]=(int16_t)(buffer[i*2] | (buffer[i*2+1]<<8));
                din[i][1]=0;
          }
          for(int i=buffer_frames;i<N;i++)
          {
              din[i][0]=0;
              din[i][1]=0;
          }

          fftw_execute(p);

          double maxAmpl=0;
          int freq=-1;
          int startFreqIndex=1100/rate;
          int endFreqIndex=startFreqIndex+100/rate*64;
          int delta=100/rate;
          for(int i=startFreqIndex;i<endFreqIndex;i+=delta)
          {
                double ampl=(out[i][0]*out[i][0]+out[i][1]*out[i][1])/N;
                if(ampl>maxAmpl)
                {
                    freq=i*rate;
                    maxAmpl=ampl;
                }
          }

          if(this->frameLength==0)
          {
              if(maxAmpl>threshold && (freq==freqS0))
              {
                  this->frameLength=1;
              }
          }
          else if(this->frameLength==1)
          {
                if(maxAmpl>threshold)
                {
                    if(freq==freqS1)
                    {
                        this->frameLength=2;
                    }
                    else if(freq==freqS0)
                    {
                        this->frameLength=1;
                    }
                    else
                    {
                        this->frameLength=0;
                    }
                }
                else
                {
                    this->frameLength=0;
                }
          }
          else
          {
                if(maxAmpl>threshold/2)
                {

                    if((freq==freqS0) || (freq==freqS1))
                    {
                        if(this->frameLength>2)
                        {
                            std::string txt(this->frameTxt+2,this->frameLength-2);
                            std::cout<<txt<<std::endl;
                        }
                        this->frameLength=0;
                    }
                    else
                    {

                        int codeIndex=(freq-1100)/100-1;
                        if(codeIndex>=0 && codeIndex<=9)
                        {
                            this->frameTxt[frameLength]=(char)('0'+codeIndex);
                        }
                        else if(codeIndex>=10 && codeIndex<10+26)
                        {
                            this->frameTxt[frameLength]=(char)('A'+codeIndex-10);
                        }
                        else if(codeIndex>=10+26 && codeIndex<10+26+26)
                        {
                            this->frameTxt[frameLength]=(char)('a'+codeIndex-10-26);
                        }

                        frameLength++;

                        if(frameLength>=TxtSize)
                        {
                            frameLength=0;
                        }

                    }


                }
                else
                {
                    if(this->frameLength>2)
                    {
                        std::string txt(this->frameTxt+2,this->frameLength-2);
                        std::cout<<txt<<std::endl;
                    }

                    this->frameLength=0;
                }
          }
          if(maxAmpl>100)
          {
            //std::cout<<freq<<","<<maxAmpl<<std::endl;
          }

      }

      free(buffer);

      fprintf(stdout, "buffer freed\n");


      snd_pcm_close (capture_handle);

      fftw_destroy_plan(p);
      fftw_cleanup();


      if(din!=NULL) fftw_free(din);
      if(out!=NULL) fftw_free(out);


      std::cout<<"audio interface closed\n"<<std::endl;
}


void SoundDecode::start()
{

}

void SoundDecode::stop()
{

}
