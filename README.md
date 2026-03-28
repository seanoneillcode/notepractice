# TODO

# DONE

hide/show guide
proper chaaracters for buttons
add icon
remove quit button
line guide out of stave notes
only score one point per touch
correct note

## deploying the app

```
cd wrapper
export PATH=/Applications/Android\ Studio.app/Contents/jbr/Contents/Home/bin:$PATH
export ANDROID_HOME=~/Library/Android/sdk
ebitenmobile bind -target android -javapkg com.seanoneillcode.mobiletest -o mobiletest.aar -androidapi 35 .
cp mobiletest.aar ~/AndroidStudioProjects/MobileTest/app/lib
```
