#include "../src/RtcTokenBuilder.h"
#include <gtest/gtest.h>
#include <string>
#include <stdint.h>
using namespace agora::tools;

class RtcTokenBuilder_test : public testing::Test {
 protected:
  virtual void SetUp() {
    appID = "972CA35de60c44645bbae8a215061b33";
    appCertificate = "5CFd2fd1756d40ecb72977518be15d3b";
    userAccount = "test_user";
    crcUserAccount = crc32(
        0, reinterpret_cast<Bytef*>(const_cast<char*>(userAccount.c_str())),
        userAccount.length());
    cname = "test_channel";
    crcCname = crc32(
        0, reinterpret_cast<Bytef*>(const_cast<char*>(cname.c_str())),
        cname.length());

  }
  void testRtcTokenBuilder();
  virtual void TearDown() {}


 private:
  std::string appID;
  std::string appCertificate;
  std::string userAccount;
  uint32_t crcUserAccount;
  std::string cname;
  uint32_t crcCname;
};

void RtcTokenBuilder_test::testRtcTokenBuilder() {
    uint32_t joints = 1614049514;
    uint32_t audiots = 1614049515;
    uint32_t videots = 1614049516;
    uint32_t datats = 1614049517;
    std::string token = RtcTokenBuilder::buildTokenWithUserAccount(
        appID, appCertificate, cname, userAccount,
        joints, audiots, videots, datats);
    AccessToken parser;
    parser.FromString(token);
    EXPECT_EQ(parser.app_id_, appID);
    EXPECT_EQ(parser.crc_channel_name_ , crcCname);
    EXPECT_EQ(parser.crc_uid_, crcUserAccount);
    EXPECT_EQ(parser.message_.messages[AccessToken::kJoinChannel],
        joints);
    EXPECT_EQ(parser.message_.messages[AccessToken::kPublishAudioStream],
        audiots);
    EXPECT_EQ(parser.message_.messages[AccessToken::kPublishVideoStream],
        videots);
    EXPECT_EQ(parser.message_.messages[AccessToken::kPublishDataStream],
        datats);
    EXPECT_EQ(
        AccessToken::GenerateSignature(
          appCertificate, appID, cname, userAccount,
          parser.message_raw_content_),
        parser.signature_);
}

TEST_F(RtcTokenBuilder_test, testRtcTokenBuilder) {
  testRtcTokenBuilder();
}

