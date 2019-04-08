#include <QCoreApplication>
#include <QImage>
 
int main(int argc, char *argv[])
{
    QCoreApplication a(argc, argv);
    Q_UNUSED(a)
 
    QImage inImage("fred.jpg");
    inImage = inImage.convertToFormat(QImage::Format_RGB32);
    QImage outImage(inImage.size(), inImage.format());
 
    int kw = 3;
    int kh = 3;
    qreal kernel[] = {0, 0, 0,
                      0, 1, 0,
                      0, 0, 0};
 
    int offsetX = (kw - 1) / 2;
    int offsetY = (kh - 1) / 2;
 
    for (int y = 0; y < inImage.height(); y++) {
        QRgb *outLine = (QRgb *) outImage.scanLine(y);
 
        for (int x = 0; x < inImage.width(); x++) {
            qreal pixelR = 0;
            qreal pixelG = 0;
            qreal pixelB = 0;
 
            // Apply convolution to each channel.
            for (int j = 0; j < kh; j++) {
                if (y + j < offsetY
                    || y + j - offsetY >= inImage.height())
                    continue;
 
                const QRgb *inLine = (QRgb *) inImage.constScanLine(y + j - offsetY);
 
                for (int i = 0; i < kw; i++) {
                    if (x + i < offsetX
                        || x + i - offsetX >= inImage.width())
                        continue;
 
                    qreal k = kernel[i + j * kw];
                    QRgb pixel = inLine[x + i - offsetX];
 
                    pixelR += k * qRed(pixel);
                    pixelG += k * qGreen(pixel);
                    pixelB += k * qBlue(pixel);
                }
            }
 
            quint8 r = qBound(0., pixelR, 255.);
            quint8 g = qBound(0., pixelG, 255.);
            quint8 b = qBound(0., pixelB, 255.);
            outLine[x] = qRgb(r, g, b);
        }
    }
 
    outImage.save("out.png");
 
    return EXIT_SUCCESS;
}