package constant

const ChatbotSystemContent = "Role: Kamu adalah seorang ahli di bidang zero waste di Indonesia dengan pengalaman selama 30 tahun. Kamu memiliki pengetahuan mendalam tentang perusahaan, komunitas, startup, lembaga, dan acara-acara yang berhubungan dengan zero waste di Indonesia.\n Context: Redooce Hub adalah aplikasi yang digunakan sebagai wadah untuk organisasi, lembaga, atau individu untuk berkolaborasi dalam berbagai kegiatan yang berfokus pada penanganan masalah sampah atau zero waste.\n Responsibility: Kamu akan menjadi penasehat bagi pengguna yang ingin mengajak kerjasama organisasi lain di bidang zero waste.\n Scope: Kamu hanya akan menjawab pertanyaan yang berhubungan dengan zero waste. Jika ada pertanyaan yang tidak berhubungan dengan zero waste, kamu akan selalu berusaha menghubungkannya ke bidang zero waste"

//http error
var ErrBadRequest = "Bad Request"
var ErrValidation = "make sure you follow the input requirements"
var ErrNotFound = "Not Found"
var ErrInternalServer = "Internal Server Error"

//address error message
var ErrCreateUserAddress = "Failed to create user address"
var ErrCreateOrganizationAddress = "Failed to create organization address"
var ErrParameterNotFound = "make sure you input the parameter organization_id or user_id"
var ErrDeleteAddress = "Failed to delete address"

// http success
var SuccessCreated = "Created"
var SuccessOk = "Ok"

// address success message
var SuccessCreateAddress = "Success created new address"
var SuccessDeleteAddress = "Success delete address"

// collaboration error message
var ErrCreateCollaboration = "Failed to create collaboration"
var ErrNotFoundCollaboration = "Collaboration not found"
var ErrFailedGetFile = "Failed to get file"
var ErrFailedOpenFile = "Failed to open file"
var ErrUploadFile = "Failed to upload file"
var ErrFailedProposal = "Failed to create proposal"
var ErrFailedSendEmail = "Failed to send email"
var ErrFailedDeleteCollaboration = "Failed to delete collaboration"

//collaboration success message
var SuccessCreateCollaboration = "Success created new collaboration"
var SuccessGetCollaboration = "Success get collaboration"
var SuccessUpdateCollaboration = "Success update collaboration"
var SuccessGetAllCollaboration = "Success get all collaboration"
var SuccessDeleteCollaboration = "Success delete collaboration"