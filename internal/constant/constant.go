package constant

const ChatbotSystemContent = "Role: Kamu adalah seorang ahli di bidang zero waste di Indonesia dengan pengalaman selama 30 tahun. Kamu memiliki pengetahuan mendalam tentang perusahaan, komunitas, startup, lembaga, dan acara-acara yang berhubungan dengan zero waste di Indonesia.\n Context: Redooce Hub adalah aplikasi yang digunakan sebagai wadah untuk organisasi, lembaga, atau individu untuk berkolaborasi dalam berbagai kegiatan yang berfokus pada penanganan masalah sampah atau zero waste.\n Responsibility: Kamu akan menjadi penasehat bagi pengguna yang ingin mengajak kerjasama organisasi lain di bidang zero waste.\n Scope: Kamu hanya akan menjawab pertanyaan yang berhubungan dengan zero waste. Jika ada pertanyaan yang tidak berhubungan dengan zero waste, kamu akan selalu berusaha menghubungkannya ke bidang zero waste"

//http error
var ErrBadRequest = "Bad Request"
var ErrValidation = "make sure you follow the input requirements"
var ErrNotFound = "Not Found"
var ErrInternalServer = "Internal Server Error"
var ErrUnauthorized = "Unauthorized"
var ErrInvalidCredentials = "Invalid credentials"

//error binding
var ErrBinding = "Error while binding your request"

// http success
var SuccessCreated = "Created"
var SuccessOk = "Ok"
var Conflict = "Conflict"

//file success message
var ErrFailedGetFile = "Failed to get file"
var ErrFailedOpenFile = "Failed to open file"
var ErrFailedUploadFile = "Failed to upload file"

//address error message
var ErrCreateUserAddress = "Failed to create user address"
var ErrCreateOrganizationAddress = "Failed to create organization address"
var ErrParameterNotFound = "make sure you input the parameter organization_id or user_id"
var ErrDeleteAddress = "Failed to delete address"
var ErrGetAllAddress = "Failed to get all address by organization_id or user_id"

// address success message
var SuccessCreateAddress = "Success created new address"
var SuccessDeleteAddress = "Success delete address"
var SuccessGetAllAddress = "Success get all address by organization_id or user_id"

// collaboration error message
var ErrGetCollaboration = "Failed to get collaboration"
var ErrCreateCollaboration = "Failed to create collaboration"
var ErrNotFoundCollaboration = "Collaboration not found"
var ErrSendEmail = "Failed to send email"
var ErrCreateProposal = "Failed to create proposal"
var ErrDeleteCollaboration = "Failed to delete collaboration"
var ErrUpdateCollaboration = "Failed to update collaboration"
var ErrGetAllCollaboration = "Failed to get all collaboration"
var ErrDeleteProposal = "Failed to delete proposal"

//collaboration success message
var SuccessCreateCollaboration = "Success created new collaboration"
var SuccessGetCollaboration = "Success get collaboration"
var SuccessUpdateCollaboration = "Success update collaboration"
var SuccessGetAllCollaboration = "Success get all collaboration"
var SuccessDeleteCollaboration = "Success delete collaboration"

// organization error message
var ErrGetAllOrganization = "Failed to get all organization"
var ErrGetOrganization = "Failed to get organization"
var ErrNotFoundOrganization = "Organization not found"
var ErrCreateOrganization = "Failed to create organization"
var ErrUpdateOrganization = "Failed to update organization"
var ErrDeleteOrganization = "Failed to delete organization"
var ErrGetOrganizationByUserId = "Failed to get organization by user"

// organization success message
var SuccessGetAllOrganization = "Success get all organization"
var SuccessGetOrganization = "Success get organization"
var SuccessCreateOrganization = "Success create organization"
var SuccessUpdateOrganization = "Success update organization"
var SuccessDeleteOrganization = "Success delete organization"
var SuccessGetOrganizationByUserId = "Success get organization by user"

// user error message
var ErrGetUser = "Failed to get user"
var ErrGetProfile = "Failed to get profile"
var ErrRegisterUser = "Failed to register user"
var ErrAlreadyRegistered = "User already registered"
var ErrEmailAlreadyExist = "Email already exist"
var ErrGeneratePassword = "Unknown error when generate password"
var ErrAccessToken = "Failed to generate access token"
var ErrRefreshToken = "Failed to generate refresh token"
var ErrNotFoundUser = "User not found"
var ErrUpdateUser = "Failed to update user"
var ErrGetDashboardData = "Failed to get dashboard data"
var ErrSendMessageChatbot = "Failed to send message chatbot"


// user success message
var SuccessGetUser = "Success get user"
var SuccessRegisterUser = "Success register user"
var SuccessLoginUser = "Success login"
var SuccessRefreshToken = "Success refresh token"
var SuccessUpdateUser = "Success update user"
var SuccessGetDashboardData = "Success get dashboard data"
var SuccessSendMessageChatbot = "Success send message chatbot"

//auth error message
var ErrFailedExtractToken = "Failed to extract user ID from token"
var ErrUnauthorizedAuth = "You are not authorized to access this resource"